//go:build wireinject
// +build wireinject

package main

import (
	"github.com/tunangoo/full-time-go-dev/internal/handler"
	"github.com/tunangoo/full-time-go-dev/internal/repository"
	"github.com/tunangoo/full-time-go-dev/internal/service"

	"github.com/google/wire"
	"github.com/uptrace/bun"
)

func wireApp(db *bun.DB) (*handler.Handler, error) {
	panic(wire.Build(handler.ProviderSet, service.ProviderSet, repository.ProviderSet))
}
