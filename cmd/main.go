package main

import (
	fyneapp "fyne.io/fyne/v2/app"
	application "tutor-app/internal/app"
	"tutor-app/internal/gui"
)

func main() {
	coreApp := fyneapp.New()
	appInstance := application.NewApp()

	gui.ShowMainWindow(appInstance, coreApp)
}
