package loms

import (
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/domain"
	"net/http"
)

type Module struct {
}

func InitModule(logger domain.Logger) {

}

func (m *Module) RegisterHTTPHandler(handler http.Handler) {

}
