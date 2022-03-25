package main

import (
	"log"

	"github.com/maddiehayes/ctxman/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
