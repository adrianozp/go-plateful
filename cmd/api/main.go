package main

import (
	"context"

	"github.com/adrianozp/go-plateful/cmd/api/modules"
	"github.com/adrianozp/go-plateful/internal/db"
	"github.com/adrianozp/go-plateful/pkg/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	factories := fx.Provide(
		config.LoadConfig,
		gin.Default,
		db.NewGormDB,
	)

	options := fx.Options(
		factories,

		modules.PlaceInjections(),
		modules.PlaceFactories(),
		modules.PlaceEndpoints(),

		modules.ReviewInjections(),
		modules.ReviewFactories(),
		modules.ReviewEndpoints(),

		fx.Invoke(func(lc fx.Lifecycle, server *gin.Engine) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go server.Run()
					return nil
				},
			})
		}),
	)
	app := fx.New(options)
	app.Run()
}
