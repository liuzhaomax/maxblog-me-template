package main

import (
	"context"
	"flag"
	"maxblog-me-template/internal/app"
)

func main() {
	config := flag.String("c", "env/dev.yaml", "配置文件")
	flag.Parse()
	ctx := context.Background()
	app.Launch(
		ctx,
		app.SetConfigFile(*config),
	)
}
