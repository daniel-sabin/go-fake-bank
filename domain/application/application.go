package application

import (
	"fmt"

	"github.com/google/uuid"
)

type Application struct {
	Name         string
	ClientId     uuid.UUID
	ClientSecret uuid.UUID
}

type ApplicationStore interface {
	save(app *Application)
	readApplication(name string) *Application
}

func ApplicationStoreFactory(ar ApplicationStore) func(name string) *Application {
	return func(name string) *Application {
		application := Application{
			Name:         name,
			ClientId:     uuid.New(),
			ClientSecret: uuid.New(),
		}

		ar.save(&application)

		return &application
	}
}

func ApplicationReadStoreFactory(ar ApplicationStore) func(name string) *Application {
	return func(name string) *Application {
		app := ar.readApplication(name)

		if app == nil {
			fmt.Printf("no application found with name %v in store", name)
			return nil
		}

		return app
	}
}
