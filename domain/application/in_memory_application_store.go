package application

import (
	"fmt"
	"sync"
)

type InMemoryApplicationStore struct {
	store sync.Map
}

func NewInMemoryApplicationStore() *InMemoryApplicationStore {
	return &InMemoryApplicationStore{sync.Map{}}
}

func (i *InMemoryApplicationStore) save(app *Application) {
	fmt.Printf("trying to save application %v\r\n", app.Name)
	i.store.Store(app.Name, app)
}

func (i *InMemoryApplicationStore) readApplication(name string) *Application {
	app, ok := i.store.Load(name)

	if !ok || app == nil {
		fmt.Printf("no application found with name %v\r\n", name)
		return nil
	}

	return app.(*Application)
}
