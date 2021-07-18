package foo

import (
	"flag"
	"log"
	"os"
	"yaml-2"
)

type Config struct {
	Port         string
	Db_url       string
	Jaeger_url   string
	Sentry_url   string
	Kafka_broker string
	Some_app_id  int
	Some_app_key int
}

func FileConfig() *Config {

	d_conf := Config{}

	var path_config = flag.String("path_config", "none", "path to file config")
	flag.Parse()
	file, err := os.Open(*path_config)
	if err != nil {
		log.Printf("Не могу открыть файл: %v", err)
	}

	f_info, err := os.Stat(*path_config)
	if err != nil {
		log.Fatalln(err)
	}

	data := make([]byte, f_info.Size())

	_, err = file.Read(data)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(data)
	err = yaml.Unmarshal([]byte(data), &d_conf)
	//err = yaml.NewDecoder(file).Decode(&d_conf)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	return &d_conf
}

func ConsConfig() *Config {
	var port_cache = flag.String("port", "8080", "Port number")
	var db_url_cache = flag.String("db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "Database url")
	var jaeger_url_cache = flag.String("jaeger_url", "http://jaeger:16686", "jaeger  link")
	var sentry_url_cache = flag.String("sentry_url", "http://sentry:9000", "sentry link")
	var kafka_broker_cache = flag.String("kafka_broker", "kafka:9092", "user info")
	var some_app_id_cache = flag.Int("id", 0, "Some app id")
	var some_app_key_cache = flag.Int("key", 0, "Some app key")

	flag.Parse()

	return &Config{
		Port:         *port_cache,
		Db_url:       *db_url_cache,
		Jaeger_url:   *jaeger_url_cache,
		Sentry_url:   *sentry_url_cache,
		Kafka_broker: *kafka_broker_cache,
		Some_app_id:  *some_app_id_cache,
		Some_app_key: *some_app_key_cache,
	}
}
