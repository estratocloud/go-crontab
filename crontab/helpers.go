package crontab

import (
	hourfield "github.com/estratocloud/go-crontab/crontab/hour"
	minutefield "github.com/estratocloud/go-crontab/crontab/minute"
)

// EveryXMinutes Only truly "every" if the minutes divides 60 (eg 1,2,3,4,5,6,10,12,15,20,30)
// However this function will generate an expression that is pretty close for other numbers,
// eg EveryXMinutes(25) would run on 0, 25, 50 then again 10 minutes later on the next 0.
// Therefor passing minutes not between 1-30 will produce an error to avoid confusing results.
func EveryXMinutes(minutes int) *Expression {
	if minutes == 1 {
		return NewExpression(nil)
	}

	return NewExpression(&Expression{
		Minute: minutefield.EveryX(minutes),
	})
}

// EveryXHours Only truly "every" if the hours divides 24 (eg 1,2,3,4,6,8,12)
// However this function will generate an expression that is pretty close for other numbers,
// eg EveryXHours(7) would run on 0, 7, 14, 21 and then again 3 hours later on the next 0.
// Therefor passing hours not between 1-12 will produce an error to avoid confusing results.
func EveryXHours(hours int) *Expression {
	if hours == 1 {
		return NewExpression(&Expression{
			Minute: minutefield.Minute(0),
		})
	}

	return NewExpression(&Expression{
		Minute: minutefield.Minute(0),
		Hour:   hourfield.EveryX(hours),
	})
}

// EveryHour Generate an expression that will run once an hour, on the specified minute
func EveryHour(minute int) *Expression {
	return NewExpression(&Expression{
		Minute: minutefield.Minute(minute),
	})
}

// EveryDayAt Generate an expression that will run once a day, on the specified hour and minute
func EveryDayAt(hour int, minute int) *Expression {
	return NewExpression(&Expression{
		Minute: minutefield.Minute(minute),
		Hour:   hourfield.Hour(hour),
	})
}
