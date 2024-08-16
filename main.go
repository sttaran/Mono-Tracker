package main

import (
	"embed"
	"github.com/playwright-community/playwright-go"
	appModule "mono-tracker/internal/app"
	"mono-tracker/pkg"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	browsers := []string{"chromium"}
	err := playwright.Install(&playwright.RunOptions{Browsers: browsers})
	const dbPath = "./mono-tracker"
	dbClient := pkg.NewSQLiteClient(&pkg.Config{ConnectionURL: dbPath})
	err = dbClient.Open()
	if err != nil {
		return
	}
	// Create an instance of the app structure
	app := appModule.NewApp(dbClient)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "mono-tracker",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
