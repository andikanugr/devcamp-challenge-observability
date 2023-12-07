//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"todoapp/config"
	"todoapp/internal/delivery/http"
)

func InitHTTPServer(*config.Schema) *http.Server {
	wire.Build(httpSet)
	return &http.Server{}
}
