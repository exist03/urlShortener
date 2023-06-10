package main

import (
	"context"
	"ozon/config"
	"ozon/internal/app"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfigEnv()
	app.Run(ctx, cfg)
}
