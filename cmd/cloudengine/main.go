// cmd/cloudengine/main.go
package main

import (
	"flag"
	"log"
	"os"

	"cloudengine/internal/cloudengine"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := cloudengine.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
