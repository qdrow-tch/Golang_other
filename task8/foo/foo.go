package foo

import (
	"flag"
)

type Config struct {
	port         string
	db_url       string
	some_app_id  int
	some_app_key int
}

func NewConfig() *Config {
	var port_cache = flag.String("port", "8080", "Port number")
	var db_url_cache = flag.String("db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "Database url")
	var some_app_id_cache = flag.Int("id", 0, "Some app id")
	var some_app_key_cache = flag.Int("key", 0, "Some app key")

	flag.Parse()

	return &Config{
		port:         *port_cache,
		db_url:       *db_url_cache,
		some_app_id:  *some_app_id_cache,
		some_app_key: *some_app_key_cache,
	}
}
