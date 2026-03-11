package main

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/xjh22222228/open-pos/server"
)

//go:embed all:frontend/dist/client
var assets embed.FS

func main() {
	go func() {
		server.Run()
	}()

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "OPEN POS",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			app,
		},

		Debug: options.Debug{
			OpenInspectorOnStartup: true, // 启动时自动打开开发者工具
		},

		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			runtime.WindowMaximise(ctx)
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
