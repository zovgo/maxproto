package protocol

import (
	"strconv"
	"time"
)

type Timestamp time.Time

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	seconds, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = Timestamp(time.Unix(seconds, 0))
	return nil
}
