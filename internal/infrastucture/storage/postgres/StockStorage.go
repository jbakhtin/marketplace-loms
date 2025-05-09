package postgres

type StockStorage struct {
}

func NewStockStorage() (OrderStorage, error) {
	return OrderStorage{}, nil
}
