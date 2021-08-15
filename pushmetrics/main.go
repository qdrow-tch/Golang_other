package main

import (
	"log"
	"net/http"

	"github.com/qdrow-tch/Golang/tree/pushmetrics/pushmetrics/fileconverter"
	"github.com/qdrow-tch/Golang/tree/pushmetrics/pushmetrics/metricscreater"
)

func main() {

	exrate, err := fileconverter.ConvertToER("./exchangerate.yaml")
	if err != nil {
		log.Fatalln("невозможно конвертировать в ExchangeRate: ", err)
	}

	http.HandleFunc("/metrics", metricscreater.GetERmetrics(*exrate))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Не удалось запустить сервер", err)
	}

}
