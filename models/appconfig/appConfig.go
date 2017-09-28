package appconfig

//AppConfig ...
type AppConfig struct {
	Envs []Environment
}

//Environment ...
type Environment struct {
	Name     string
	OAuthURL string
	BaseURL  string
	Tenant   string
	UserName string
	Password string
}
