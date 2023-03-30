package ports

type Client struct {
	Name          string `json:"client_name"`
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
}

type ClientsStore interface {
	Client(n string) Client
}
