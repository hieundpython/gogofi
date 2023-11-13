//go:build wireinject
// +build wireinject

package connect

import (
	"github.com/google/wire"

	"gogofi/internal/controller"
	"gogofi/internal/database/services"
)

func InitializeApp(db services.DBTX) *controller.UserController {
	wire.Build(controller.NewUserController, services.New)
	return &controller.UserController{}
}
