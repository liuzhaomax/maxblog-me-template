package main

import (
	"context"
	"maxblog-me-template/internal/app"
)

func main() {
	const ConfigDir = "env"
	const ConfigFile = "dev.yaml"
	ctx := context.Background()
	app.Launch(
		ctx,
		app.SetConfigDir(ConfigDir),
		app.SetConfigFile(ConfigFile),
	)
}
