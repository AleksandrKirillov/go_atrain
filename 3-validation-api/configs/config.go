package configs

type Config struct {
	Auth AuthConfig
}

type AuthConfig struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *Config {
	return &Config{
		Auth: AuthConfig{
			Email:    "test@test.ru",
			Password: "123",
			Address:  "test@test.ru",
		},
	}
}
