package repository

import "engineecore/demobank-server/domain/ports"

type ClientsStore struct {
	Store         map[string]ports.Client
	uuidGenerator func() string
}

func NewInMemoryClientsStore(uuidGenerator func() string) *ClientsStore {
	return &ClientsStore{make(map[string]ports.Client), uuidGenerator}
}

func (cs *ClientsStore) Client(n string) ports.Client {
	if _, exists := cs.Store[n]; !exists {
		cs.Store[n] = ports.Client{
			Name:          n,
			Client_id:     cs.uuidGenerator(),
			Client_secret: cs.uuidGenerator(),
		}
	}
	return cs.Store[n]
}
