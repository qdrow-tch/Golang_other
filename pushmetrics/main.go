package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"yaml-2"

	"github.com/qdrow-tch/Golang/tree/pushmetrics/pushmetrics/fileconverter"
	"github.com/qdrow-tch/Golang/tree/pushmetrics/pushmetrics/metricscreater"
)

type Config struct {
	FileConvert string
	Port        string
}

func main() {

	var config Config
	yaml_config, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalln("Невозможно открыть файл конфигурации")
	}
	err = yaml.Unmarshal(yaml_config, &config)
	if err != nil {
		log.Fatalln("Невозможно преобразовать из .yaml в Config")
	}

	exrate, err := fileconverter.ConvertToER(config.FileConvert)
	if err != nil {
		log.Fatalln("невозможно конвертировать в ExchangeRate: ", err)
	}

	http.HandleFunc("/metrics", metricscreater.GetERmetrics(*exrate))
	err = http.ListenAndServe(config.Port, nil)
	if err != nil {
		log.Fatalln("Не удалось запустить сервер", err)
	}

}
