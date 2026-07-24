# go-crontab
A small library for generating crontab expressions

Full documentation is available at https://pkg.go.dev/github.com/estratocloud/go-crontab

[![release](https://img.shields.io/badge/dynamic/json?url=https://proxy.golang.org/github.com/estratocloud/go-crontab/@latest&query=$.Version&label=version)](https://pkg.go.dev/github.com/estratocloud/go-crontab?tab=versions)
[![build](https://github.com/estratocloud/go-crontab/actions/workflows/buildcheck.yaml/badge.svg)](https://github.com/estratocloud/go-crontab/actions/workflows/buildcheck.yaml?query=branch%3Amain)


### Installation
```sh
go get github.com/estratocloud/go-crontab
```


### Getting Started
The main structure used by this library is the `crontab.Expression`, and the `NewExpression()` function is the main entry point:
```go
import (
	"github.com/estratocloud/go-crontab/crontab"
	"github.com/estratocloud/go-crontab/crontab/day"
	"github.com/estratocloud/go-crontab/crontab/hour"
	"github.com/estratocloud/go-crontab/crontab/minute"
	"github.com/estratocloud/go-crontab/crontab/month"
	"github.com/estratocloud/go-crontab/crontab/weekday"
)
e := crontab.NewExpression(&crontab.Expression{
	Minute: minute.Every(),
	Hour: hour.Every(),
	Day: day.Every(),
	Month: month.Every(),
	Weekday: weekday.Every(),
})
println(e.String()) // "* * * * *"
```

All of the fields are optional, including the initial structure itself:
```go
e := crontab.NewExpression(&crontab.Expression{
	Day: day.Day(23),
})
println(e.String()) // "* * 23 * *"

e := crontab.NewExpression(nil)
println(e.String()) // "* * * * *"
```

When generating fields you should check for errors if passing user input:
```go
fieldMinutes := minute.EveryX(minutesFromUser)
if field.Error() != nil {
	handleError(err)
}
```

However the library does defer errors until later, which is useful if you're dealing with constant values you trust and don't need to error check every field generation:
```go
e := crontab.NewExpression(&Expression{
	Minute: minute.Minute(0),
	Hour:   hour.List(8, 12, 16),
	Weekday: weekday.Range(time.Monday, time.Friday),
})
println(e.MustFormat()) // "0 8,12,16 * * 1-5"
```

There are a few helper functions available for common expressions:
```go
import "github.com/estratocloud/go-crontab/crontab"

line := fmt.Sprintf("%s root myjob", crontab.EveryXMinutes(10).MustFormat()) // "*/10 * * * *"
line := fmt.Sprintf("%s root myjob", crontab.EveryXHours(4).MustFormat()) // "0 */4 * * *"
line := fmt.Sprintf("%s root myjob", crontab.EveryHour(45).MustFormat()) // "45 * * * *"
line := fmt.Sprintf("%s root myjob", crontab.EveryDayAt(7, 30).MustFormat()) // "30 7 * * *"
```

### Where to get help
Found a bug? Got a question? Just not sure how something works?  
Please [create an issue](https://github.com/estratocloud/go-crontab/issues) and we'll do our best to help out.  
Alternatively you can connect with us on [LinkedIn](https://www.linkedin.com/company/estratocloud/)
