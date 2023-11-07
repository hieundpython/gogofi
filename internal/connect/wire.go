//go:build wireinject
// +build wireinject

package connect

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"

	"gogofi/internal/handler"
	"gogofi/internal/repo"
)

func InitializeApp(db *sqlx.DB) *handler.UserHandler {
	wire.Build(handler.NewUserHandler, repo.NewRepo)
	return &handler.UserHandler{}
}
