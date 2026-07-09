package day

import (
	"github.com/estratocloud/go-crontab/crontab/internal"
)

var Every = internal.Every

func Day(value int) internal.Field {
	return internal.Value("day", 1, 31, value)
}

func List(values ...int) internal.Field {
	return internal.List("day", 1, 31, values...)
}

func Range(from int, to int) internal.Field {
	return internal.Range("day", 1, 31, from, to)
}

func EveryX(value int) internal.Field {
	return internal.EveryX("day", 15, value)
}
