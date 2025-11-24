package postgres

import (
	"database/sql"
	"io/fs"

	"github.com/pressly/goose/v3"

	_ "github.com/jackc/pgx/v5/stdlib"
	orderpostgres "github.com/jbakhtin/marketplace-loms/pkg/order/infra/postgres"
	stockpostgres "github.com/jbakhtin/marketplace-loms/pkg/stock/infra/postgres"
	"github.com/pkg/errors"
)

type Config interface {
	GetDbDriver() string
	GetDbHost() string
	GetDbPort() string
	GetDbName() string
	GetDbUser() string
	GetDbPassword() string
}

// mergedFS объединяет миграции из всех модулей в один FS
type mergedFS struct {
	filesystems []fs.FS
}

func (m *mergedFS) Open(name string) (fs.File, error) {
	// Пробуем открыть файл из каждого FS по очереди
	for _, fsys := range m.filesystems {
		if file, err := fsys.Open(name); err == nil {
			return file, nil
		}
	}
	return nil, fs.ErrNotExist
}

func (m *mergedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	// Собираем все файлы из всех FS
	allEntries := make(map[string]fs.DirEntry)

	for _, fsys := range m.filesystems {
		entries, err := fs.ReadDir(fsys, name)
		if err != nil {
			continue // Пропускаем если директория не найдена в этом FS
		}
		for _, entry := range entries {
			allEntries[entry.Name()] = entry
		}
	}

	// Преобразуем map в slice
	result := make([]fs.DirEntry, 0, len(allEntries))
	for _, entry := range allEntries {
		result = append(result, entry)
	}
	return result, nil
}

func (m *mergedFS) ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(m, name)
}

func NewSQLClient(cfg Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.GetDbDriver(), getDbDSN(cfg))
	if err != nil {
		return nil, errors.Wrap(err, "db open")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "db ping")
	}

	err = goose.SetDialect("postgres")
	if err != nil {
		return nil, errors.Wrap(err, "set dialect")
	}

	// Объединяем все миграции из всех модулей в один FS
	// Используем fs.Sub для получения поддиректории migrations из каждого модуля
	stockMigrations, err := fs.Sub(stockpostgres.Migrations, "migrations")
	if err != nil {
		return nil, errors.Wrap(err, "get stock migrations")
	}

	orderMigrations, err := fs.Sub(orderpostgres.Migrations, "migrations")
	if err != nil {
		return nil, errors.Wrap(err, "get order migrations")
	}

	// Создаем объединенный FS
	combinedFS := &mergedFS{
		filesystems: []fs.FS{stockMigrations, orderMigrations},
	}

	// Запускаем все миграции из объединенного FS
	// Goose сам отсортирует их по версиям и применит в правильном порядке
	goose.SetBaseFS(combinedFS)
	err = goose.Up(db, ".")
	if err != nil {
		return nil, errors.Wrap(err, "run migrations")
	}

	return db, nil
}

func getDbDSN(cfg Config) string {
	return "postgres://" + cfg.GetDbUser() + ":" + cfg.GetDbPassword() + "@" + cfg.GetDbHost() + ":" + cfg.GetDbPort() + "/" + cfg.GetDbName()
}
