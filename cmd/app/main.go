package main

import (
	"context"
	"ozon/config"
	"ozon/internal/app"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfigEnv()
	a := app.New(ctx, cfg)
	app.Run(a)
}
