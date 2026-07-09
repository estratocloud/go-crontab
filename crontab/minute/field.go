package minute

import (
	"github.com/estratocloud/go-crontab/crontab/internal"
)

var Every = internal.Every

func Minute(value int) internal.Field {
	return internal.Value("minute", 0, 59, value)
}

func List(values ...int) internal.Field {
	return internal.List("minute", 0, 59, values...)
}

func Range(from int, to int) internal.Field {
	return internal.Range("minute", 0, 59, from, to)
}

func EveryX(value int) internal.Field {
	return internal.EveryX("minute", 30, value)
}
