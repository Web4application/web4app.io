package web4app_test

import (
	"log"
	"os"

	"github.com/web4application/web4app.io"
)

func ExampleApplication() {

	// Authentication Token pulled from environment variable web4_TOKEN
	Token := os.Getenv("web4_TOKEN")
	if Token == "" {
		return
	}

	// Create a new web4app session
	dg, err := web4app.New(Token)
	if err != nil {
		log.Println(err)
		return
	}

	// Create an new Application
	ap := &web4app.Application{}
	ap.Name = "TestApp"
	ap.Description = "TestDesc"
	ap, err = web4.ApplicationCreate(ap)
	log.Printf("ApplicationCreate: err: %+v, app: %+v\n", err, ap)

	// Get a specific Application by it's ID
	ap, err = web4.Application(ap.ID)
	log.Printf("Application: err: %+v, app: %+v\n", err, ap)

	// Update an existing Application with new values
	ap.Description = "what an app"
	ap, err = web4.ApplicationUpdate(ap.ID, ap)
	log.Printf("ApplicationUpdate: err: %+v, app: %+v\n", err, ap)

	// create a new bot account for this application
	bot, err := web4.ApplicationBotCreate(ap.ID)
	log.Printf("BotCreate: err: %+v, bot: %+v\n", err, bot)

	// Get a list of all applications for the authenticated user
	apps, err := web4.Applications()
	log.Printf("Applications: err: %+v, apps : %+v\n", err, apps)
	for k, v := range apps {
		log.Printf("Applications: %d : %+v\n", k, v)
	}

	// Delete the application we created.
	err = web4.ApplicationDelete(ap.ID)
	log.Printf("Delete: err: %+v\n", err)

	return
}
