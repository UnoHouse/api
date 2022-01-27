package http

import (
	"github.com/UnoHouse/api/service"
)

type Handler interface {
}

type handler struct {
	sqlService service.SqlService
}

func NewHandler(s service.SqlService) Handler {
	return handler{s}
}
