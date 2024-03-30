package main

import (
	"context"
	"github.com/t1ltxz-gxd/shortify/internal/app"
	"log"
)

func main() {
	// Create a new context that is never cancelled, has no values, and has no deadline.
	ctx := context.Background()

	// Call the NewApp function from the app package with the created context.
	// This function initializes a new application and returns it along with any error that might occur during the initialization.
	a, err := app.NewApp(ctx)

	// If an error occurred during the initialization of the application, log the error and terminate the program.
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	// Call the Run method of the application.
	// This method starts the application and returns any error that might occur during the execution.
	err = a.Run()

	// If an error occurred during the execution of the application, log the error and terminate the program.
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
