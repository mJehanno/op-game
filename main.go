package main

import (
	"context"
	"embed"
	"mult-game/backend/db"
	"mult-game/backend/logger"
	"mult-game/backend/score"
	"mult-game/backend/update"
	"mult-game/backend/version"
	"net/http"

	"github.com/google/go-github/v59/github"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"go.uber.org/fx"
)

//go:embed all:frontend/dist
var assets embed.FS

type Binder struct {
	//fx.In
	App     *App                    `name:"App"`
	Score   *score.ScoreService     `name:"Score"`
	Vm      *version.VersionManager `name:"VersionManager"`
	Updater *update.Updater         `name:"Updater"`
	Logger  *logger.Logger          `name:"Logger"`
	Client  *github.Client          `name:Client"`
}

type BinderParams struct {
	fx.In
	App     *App
	Score   *score.ScoreService
	Vm      *version.VersionManager
	Updater *update.Updater
	Logger  *logger.Logger
	Client  *github.Client
}

func main() {
	var bind *Binder
	ctx := func() context.Context {
		return context.Background()
	}

	fx.New(
		fx.Provide(
			update.NewErrorHandler,
			func() *github.Client {
				return github.NewClient(http.DefaultClient)
			},
			//github.NewClient,
			db.GetDbConnection,
			logger.NewLogger,
			update.NewUpdater,
			NewApp,
			score.NewScoreService,
			version.NewVersionManager,
			fx.Annotate(ctx, fx.As(new(context.Context))),
			func(p BinderParams) *Binder {
				return &Binder{
					App:     p.App,
					Score:   p.Score,
					Logger:  p.Logger,
					Vm:      p.Vm,
					Updater: p.Updater,
				}
			},
		),
		fx.Populate(&bind),
		fx.Invoke(func(b *Binder) {
			b.Updater.DoSelfUpdate(ctx())
			startApp(b)
		}),
	).Start(context.Background())
}

func startApp(bind *Binder) {
	err := wails.Run(&options.App{
		Title:  "mult-game",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        bind.App.startup,
		Bind: []interface{}{
			bind.App,
			bind.Score,
			bind.Vm,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
