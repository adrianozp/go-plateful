package modules

import (
	createHandler "github.com/adrianozp/go-plateful/app/review/handlers/create"
	findHandler "github.com/adrianozp/go-plateful/app/review/handlers/find"
	updateHandler "github.com/adrianozp/go-plateful/app/review/handlers/update"
	"github.com/adrianozp/go-plateful/app/review/repositories/mysql"

	"github.com/adrianozp/go-plateful/app/review/usecases/create"
	"github.com/adrianozp/go-plateful/app/review/usecases/find"
	"github.com/adrianozp/go-plateful/app/review/usecases/update"

	"go.uber.org/fx"
)

func ReviewFactories() fx.Option {
	return fx.Provide(
		// handlers
		createHandler.NewCreateReviewHandler,
		updateHandler.NewUpdateReviewHandler,
		findHandler.NewFindReviewHandler,

		// usecases
		find.NewFindReviewUsecase,
		create.NewCreateReviewUsecase,
		update.NewUpdateReviewUsecase,

		// repositories
		mysql.NewReviewRepository,
	)
}

func ReviewInjections() fx.Option {
	return fx.Provide(
		// usecases
		func(u find.FindReviewUsecase) findHandler.Finder { return u },
		func(u create.CreateReviewUsecase) createHandler.Creator { return u },
		func(u update.UpdateReviewUsecase) updateHandler.Updater { return u },

		// repositories
		func(r mysql.ReviewRepository) find.ReviewRepository { return r },
		func(r mysql.ReviewRepository) create.ReviewRepository { return r },
		func(r mysql.ReviewRepository) update.ReviewRepository { return r },
	)
}

func ReviewEndpoints() fx.Option {
	return fx.Module("review",
		fx.Invoke(findHandler.RegisterFindReviewRoutes),
		fx.Invoke(updateHandler.RegisterUpdateReviewRoutes),
		fx.Invoke(createHandler.RegisterCreateReviewRoutes),
	)
}
