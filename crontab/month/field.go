package month

import (
	"time"

	"github.com/estratocloud/go-crontab/crontab/internal"
)

var Every = internal.Every

func Month(value time.Month) internal.Field {
	return internal.Value("month", 1, 12, int(value))
}

func List(values ...time.Month) internal.Field {
	ints := make([]int, len(values))
	for i, v := range values {
		ints[i] = int(v)
	}
	return internal.List("month", 1, 12, ints...)
}

func Range(from time.Month, to time.Month) internal.Field {
	return internal.Range("month", 1, 12, int(from), int(to))
}

func EveryX(value time.Month) internal.Field {
	return internal.EveryX("month", 6, int(value))
}
