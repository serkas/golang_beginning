package main

import (
	"go.uber.org/fx"
	"proj/lessons/21_di/lesson/service/internal/app"
)

func main() {
	fxApp := fx.New(
		app.New(),
		fx.Invoke(app.RegisterHTTPServer),
	)

	fxApp.Run()
}
