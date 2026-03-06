package merit

import (
	"fmt"
	"time"
)

type queryDate struct {
	time.Time
	format string
}

// func (d *queryDate) UnmarshalJSON(b []byte) error {
// 	t, err := time.Parse("20060102", string(b))
// 	if err != nil {
// 		return err
// 	}
// 	d.Time = t
// 	return nil
// }

func (d queryDate) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("\"%s\"", d.Time.Format(d.format))
	return []byte(s), nil
}
