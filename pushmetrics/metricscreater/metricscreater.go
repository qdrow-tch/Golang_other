package metricscreater

import (
	"fmt"
	"net/http"

	"github.com/qdrow-tch/Golang/tree/pushmetrics/pushmetrics/fileconverter"
)

func GetERmetrics(er fileconverter.ExchangeRates) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var responseText string
		for i := range er.Currencies {
			responseText += fmt.Sprintf("exchange_rate{name=%s, value=%F}\n", er.Currencies[i].Name, er.Currencies[i].Value)
		}
		fmt.Fprintf(w, responseText)
	}
}
