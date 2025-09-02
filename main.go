package main

import (
	"embed"

	"golang-wails-reactjs/backend/application/auth"
	"golang-wails-reactjs/backend/application/student"
	"golang-wails-reactjs/backend/application/sync"
	sqlite "golang-wails-reactjs/backend/pkg/database"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	defer recover()

	db := sqlite.ConnDB()
	// Create an instance of the app structure
	app := NewApp()
	auth := auth.Provider(db)
	student := student.Provider(db)
	sync := sync.Provider(db)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "sistem-sekolah",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			auth,
			student,
			sync,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
