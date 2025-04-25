package order

type Config interface {
}

type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
	Fatal(string, ...any)
}

type Handler struct {
	cfg Config
	log Logger
}

type Item struct {
	SKU   int32
	count uint16
}

func NewOrderHandler(cfg Config, lgr Logger) (Handler, error) {
	return Handler{
		cfg: cfg,
		log: lgr,
	}, nil
}
