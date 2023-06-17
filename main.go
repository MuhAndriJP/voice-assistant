package main

import (
	"os"

	"github.com/MuhAndriJP/voice-assistant/routes"
)

func main() {
	e := routes.New()

	routes.RunPython()

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
