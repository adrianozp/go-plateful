package modules

import (
	createHandler "github.com/adrianozp/go-plateful/app/place/handlers/create"
	findHandler "github.com/adrianozp/go-plateful/app/place/handlers/find"
	updateHandler "github.com/adrianozp/go-plateful/app/place/handlers/update"
	"github.com/adrianozp/go-plateful/app/place/repositories/mysql"

	"github.com/adrianozp/go-plateful/app/place/usecases/create"
	"github.com/adrianozp/go-plateful/app/place/usecases/find"
	"github.com/adrianozp/go-plateful/app/place/usecases/update"

	"go.uber.org/fx"
)

func PlaceFactories() fx.Option {
	return fx.Provide(
		// handlers
		createHandler.NewCreatePlaceHandler,
		updateHandler.NewUpdatePlaceHandler,
		findHandler.NewFindPlaceHandler,

		// usecases
		find.NewFindPlaceUsecase,
		create.NewCreatePlaceUsecase,
		update.NewUpdatePlaceUsecase,

		// repositories
		mysql.NewPlaceRepository,
	)
}

func PlaceInjections() fx.Option {
	return fx.Provide(
		// usecases
		func(u find.FindPlaceUsecase) findHandler.Finder { return u },
		func(u create.CreatePlaceUsecase) createHandler.Creator { return u },
		func(u update.UpdatePlaceUsecase) updateHandler.Updater { return u },

		// repositories
		func(r mysql.PlaceRepository) find.PlaceRepository { return r },
		func(r mysql.PlaceRepository) create.PlaceRepository { return r },
		func(r mysql.PlaceRepository) update.PlaceRepository { return r },
	)
}

func PlaceEndpoints() fx.Option {
	return fx.Module("place",
		fx.Invoke(findHandler.RegisterFindPlaceRoutes),
		fx.Invoke(updateHandler.RegisterUpdatePlaceRoutes),
		fx.Invoke(createHandler.RegisterCreatePlaceRoutes),
	)
}
