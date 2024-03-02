package main

import (
	"context"
	"embed"
	"mult-game/backend/db"
	"mult-game/backend/logger"
	"mult-game/backend/score"
	"mult-game/backend/version"

	"github.com/blang/semver"
	selfupdate "github.com/mJehanno/ghr-self-updater"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"go.uber.org/fx"
)

//go:embed all:frontend/dist
var assets embed.FS

const (
	owner = "mJehanno"
	repo  = "op-game"
)

var (
	current = semver.MustParse("1.3.1")
)

type Binder struct {
	//fx.In
	App     *App                    `name:"App"`
	Score   *score.ScoreService     `name:"Score"`
	Vm      *version.VersionManager `name:"VersionManager"`
	Updater *selfupdate.Updater     `name:"Updater"`
	Logger  *logger.Logger          `name:"Logger"`
}

type BinderParams struct {
	fx.In
	App     *App
	Score   *score.ScoreService
	Vm      *version.VersionManager
	Updater *selfupdate.Updater
	Logger  *logger.Logger
}

func main() {
	var bind *Binder
	ctx := func() context.Context {
		return context.Background()
	}

	fx.New(
		fx.Supply(current),
		fx.Provide(
			func() *selfupdate.Updater {
				return selfupdate.New(owner, repo, current)
			},
			db.GetDbConnection,
			logger.NewLogger,
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
			bind.Updater,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
