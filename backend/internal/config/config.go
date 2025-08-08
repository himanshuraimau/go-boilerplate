package config

type Config struct{}

type Primary struct {
	Env string `koanf:"env" validate:"required"`
}
