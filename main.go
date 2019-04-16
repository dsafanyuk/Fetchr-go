package main

import (
	"fmt"

	"github.com/dsafanyuk/fetchr-go/app"
	"github.com/dsafanyuk/fetchr-go/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	fmt.Println("Running on 127.0.0.1:3000")
	app.Run(":3000")
}
