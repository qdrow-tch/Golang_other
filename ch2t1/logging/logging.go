package logging

import (
	"fmt"
	"log"
	"os"

	"github.com/qdrow-tch/errors_workout/selferrors"
)

func Push_to_logfile(message string) error {

	_, err := os.Stat("./log.txt")

	if err != nil {
		file, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_APPEND, 0644) //os.Create("./log.txt")
		if err != nil {
			return selferrors.WrapWrongRecordingError(
				fmt.Errorf("Do not create Logfile %w", err),
			)
		}
		defer func() {
			err = file.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}()

		_, err = file.Write([]byte(message + "\n"))
		if err != nil {
			return selferrors.WrapWrongRecordingError(
				fmt.Errorf("Do not create Logfile %w", err),
			)
		}

	}

	return nil
}
