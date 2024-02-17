package config

type Config struct {
	DbUrl string `env:"DB_URL,required,expand"`
	Port  int    `env:"PORT" envDefault:"3000"`
}
