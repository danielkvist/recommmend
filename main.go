package main

import (
	"log"

	"recommmended/actions"
	"recommmended/models"

	"github.com/gobuffalo/pop"
)

func main() {
	mig, err := pop.NewFileMigrator("./migrations", models.DB)
	if err != nil {
		log.Fatal(err)
	}
	mig.Up()

	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
