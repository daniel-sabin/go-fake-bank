package repository

import (
	"engineecore/demobank-server/domain/application"
	"fmt"
	"sync"
)

type inMemoryApplicationStore struct {
	store sync.Map
}

func NewInMemoryApplicationStore() application.ApplicationStore {
	return &inMemoryApplicationStore{sync.Map{}}
}

func (i *inMemoryApplicationStore) Save(app *application.Application) {
	fmt.Printf("trying to save application %v\r\n", app.Name)
	i.store.Store(app.Name, app)
}

func (i *inMemoryApplicationStore) ReadApplication(name string) *application.Application {
	app, ok := i.store.Load(name)

	if !ok || app == nil {
		fmt.Printf("no application found with name %v\r\n", name)
		return nil
	}

	return app.(*application.Application)
}
