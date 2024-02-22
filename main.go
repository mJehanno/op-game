package main

import (
	"context"
	"embed"
	"mult-game/backend/db"
	"mult-game/backend/logger"
	"mult-game/backend/score"
	"mult-game/backend/version"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"go.uber.org/fx"
)

//go:embed all:frontend/dist
var assets embed.FS

type Binder struct {
	app    *App                `name:"App"`
	score  *score.ScoreService `name:"Score"`
	vm     *version.VersionManager
	logger *logger.Logger
}

func main() {
	var bind *Binder

	fx.New(
		fx.Provide(
			NewApp,
			score.NewScoreService,
			db.GetDbConnection,
			logger.NewLogger,
			version.NewVersionManager,
			func(a *App, s *score.ScoreService, v *version.VersionManager, l *logger.Logger, d *db.Db) *Binder {
				l.DebugLogger.Debug(d)
				return &Binder{
					app:    a,
					score:  s,
					logger: l,
					vm:     v,
				}
			},
		),
		fx.Populate(&bind),
		fx.Invoke(func(b *Binder) {
			err := wails.Run(&options.App{
				Title:  "mult-game",
				Width:  1024,
				Height: 768,
				AssetServer: &assetserver.Options{
					Assets: assets,
				},
				BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
				OnStartup:        bind.app.startup,
				Bind: []interface{}{
					bind.app,
					bind.score,
					bind.vm,
				},
			})

			if err != nil {
				println("Error:", err.Error())
			}

		}),
	).Start(context.Background())
}
