package primary

import (
	"context"
)

type UseCase interface {
	CreateOrder(ctx context.Context)
	OrderInfo(ctx context.Context)
	OrderPay(ctx context.Context)
	CancelOrder(ctx context.Context)
}
