package config

type HttpServerConfig struct {
	Port string
	Host string
}

type Config struct {
	Env string
	StoragePath string
	HttpServer HttpServerConfig

}