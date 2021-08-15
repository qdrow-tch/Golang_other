package fileconverter

import (
	"log"
	"os"
	"yaml-2"
)

type ExchangeRates struct {
	Currencies []struct {
		Name  string
		Value float32
	}
}

func ConvertToER(pathtofile string) (*ExchangeRates, error) {

	d_ex := ExchangeRates{}

	file, err := os.Open(pathtofile)
	if err != nil {
		log.Printf("Ошибка открытия файла: %v", err)
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Невозможно закрыть файл: %v", err)
		}
	}()

	file_info, err := os.Stat(pathtofile)
	if err != nil {
		log.Printf("Невозможно получить информацию о файле: %v", err)
		return nil, err
	}

	data := make([]byte, file_info.Size())

	if _, err = file.Read(data); err != nil {
		log.Printf("Невозможно прочитать данные из файла: %v", err)
		return nil, err
	}

	if err := yaml.Unmarshal(data, &d_ex); err != nil {
		log.Printf("Невозможно конвертировать из файла в структуру ExchangeRates: %v", err)
		return nil, err
	}

	return &d_ex, nil

}
