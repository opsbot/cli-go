package api

// APIResponse struct
type APIResponse struct {
	Status     string
	StatusCode int
	Body       []byte
}

// Configuration struct
type Configuration struct {
	Host      string `json:"api_host"`
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

// Collection struct
type Collection struct {
	Todos []Todo `json:"spaces"`
}

// Todo struct
type Todo struct {
	Title string `json:"name"`
	Body  string `json:"description"`
}
