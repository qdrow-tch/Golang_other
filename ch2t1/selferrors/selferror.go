package selferrors

import (
	"fmt"
	"time"
)

type Wrong_recording struct {
	err    error
	timing string
}

func (er *Wrong_recording) Error() string {

	return fmt.Sprintf("Time error: %s Text error: %s", er.timing, er.err)
}

func (er *Wrong_recording) Unwrap() error {
	return er.err
}
func WrapWrongRecordingError(err_out error) error {

	return &Wrong_recording{
		err:    err_out,
		timing: time.Now().String(),
	}

}
