package domain

import "github.com/pkg/errors"

var (
	NotFoundError      = errors.New("not found")
	AlreadyExistsError = errors.New("already exists")
)
