package hour

import (
	"github.com/estratocloud/go-crontab/crontab/internal"
)

var Every = internal.Every

func Hour(value int) internal.Field {
	return internal.Value("hour", 0, 23, value)
}

func List(values ...int) internal.Field {
	return internal.List("hour", 0, 23, values...)
}

func Range(from int, to int) internal.Field {
	return internal.Range("hour", 0, 23, from, to)
}

func EveryX(value int) internal.Field {
	return internal.EveryX("hour", 12, value)
}
