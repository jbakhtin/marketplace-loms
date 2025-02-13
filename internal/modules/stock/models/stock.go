package models

type StockItemSKU int32
type StockItemCount uint64

type StockItem struct {
	SKU   StockItemSKU
	Count StockItemCount
}
