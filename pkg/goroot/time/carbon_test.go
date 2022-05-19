package time

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"testing"
)

func TestCarbon(t *testing.T) {
	fmt.Sprintf("%s", carbon.Now()) // 2020-08-05 13:14:15
	carbon.Now().ToString()         // 2020-08-05 13:14:15 +0800 CST
	carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
	// Return date of today
	carbon.Now().ToDateString() // 2020-08-05
	// Return time of today
	carbon.Now().ToTimeString() // 13:14:15
	// Return datetime of today in a given timezone
	carbon.Now(carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
	// Return timestamp with second of today
	carbon.Now().Timestamp() // 1596604455
	// Return timestamp with millisecond of today
	carbon.Now().TimestampMilli() // 1596604455000
	// Return timestamp with microsecond of today
	carbon.Now().TimestampMicro() // 1596604455000000
	// Return timestamp with nanosecond of today
	carbon.Now().TimestampNano() // 1596604455000000000

	// Return datetime of yesterday
	fmt.Sprintf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
	carbon.Yesterday().ToString()         // 2020-08-04 13:14:15 +0800 CST
	carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
	// Return date of yesterday
	carbon.Yesterday().ToDateString() // 2020-08-04
	// Return time of yesterday
	carbon.Yesterday().ToTimeString() // 13:14:15
	// Return datetime of yesterday on a given day
	carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
	// Return datetime of yesterday in a given timezone
	carbon.Yesterday(carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
	// Return timestamp with second of yesterday
	carbon.Yesterday().Timestamp() // 1596518055
	// Return timestamp with millisecond of yesterday
	carbon.Yesterday().TimestampMilli() // 1596518055000
	// Return timestamp with microsecond of yesterday
	carbon.Yesterday().TimestampMicro() // 1596518055000000
	// Return timestamp with nanosecond of yesterday
	carbon.Yesterday().TimestampNano() // 1596518055000000000
}


func TestCarbonParse(t *testing.T)  {
	carbon.Parse("2020-08-05").ToString() // 2020-08-05 00:00:00 +0800 CST
	carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("2020-08-05 13:14:15.999").ToString() // 2020-08-05 13:14:15.999 +0800 CST
	carbon.Parse("2020-08-05 13:14:15.999999").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
	carbon.Parse("2020-08-05 13:14:15.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

	carbon.Parse("2020-08-05T13:14:15+08:00").ToString() // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("2020-08-05T13:14:15.999+08:00").ToString() // 2020-08-05 13:14:15.999 +0800 CST
	carbon.Parse("2020-08-05T13:14:15.999999+08:00").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").Carbon2Time() // carbon to time

	carbon.Parse("20200805").ToString() // 2020-08-05 00:00:00 +0800 CST
	carbon.Parse("20200805131415").ToString() // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("20200805131415.999").ToString() // 2020-08-05 13:14:15.999 +0800 CST
	carbon.Parse("20200805131415.999999").ToString() // 2020-08-05 13:14:15.999999 +0800 CST
	carbon.Parse("20200805131415.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

	carbon.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
	carbon.Parse("2020-08-05 13:14:15").SetHour(24).ToDateTimeString() // 2020-08-06 00:14:15

	// Set minute
	carbon.Parse("2020-08-05 13:14:15").SetMinute(10).ToDateTimeString() // 2020-08-05 13:10:15
	carbon.Parse("2020-08-05 13:14:15").SetMinute(60).ToDateTimeString() // 2020-08-05 14:00:15

	// Set second
	carbon.Parse("2020-08-05 13:14:15").SetSecond(10).ToDateTimeString() // 2020-08-05 13:14:10
	carbon.Parse("2020-08-05 13:14:15").SetSecond(60).ToDateTimeString() // 2020-08-05 13:15:00

	// Set millisecond
	carbon.Parse("2020-08-05 13:14:15").SetMillisecond(100).Millisecond() // 100
	carbon.Parse("2020-08-05 13:14:15").SetMillisecond(999).Millisecond() // 999

	// Set microsecond
	carbon.Parse("2020-08-05 13:14:15").SetMicrosecond(100000).Microsecond() // 100000
	carbon.Parse("2020-08-05 13:14:15").SetMicrosecond(999999).Microsecond() // 999999

	// Set nanosecond
	carbon.Parse("2020-08-05 13:14:15").SetNanosecond(100000000).Nanosecond() // 100000000
	carbon.Parse("2020-08-05 13:14:15").SetNanosecond(999999999).Nanosecond() // 999999999
}
