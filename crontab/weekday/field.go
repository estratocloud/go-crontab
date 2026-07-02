package weekday

import (
	"time"

	"github.com/estratocloud/go-crontab/crontab/internal"
)

var Every = internal.Every

func Weekday(value time.Weekday) internal.Field {
	return internal.Value("weekday", 0, 6, int(value))
}

func List(values ...time.Weekday) internal.Field {
	ints := make([]int, len(values))
	for i, v := range values {
		ints[i] = int(v)
	}
	return internal.List("weekday", 0, 6, ints...)
}

func Range(from time.Weekday, to time.Weekday) internal.Field {
	return internal.Range("weekday", 0, 6, int(from), int(to))
}

func EveryX(value time.Weekday) internal.Field {
	return internal.EveryX("weekday", 3, int(value))
}
