package utils

import "time"

// TimeUnixMilli return unix milli
func TimeUnixMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// TimeFromUnixMillis return time.Time object from unix mils
func TimeFromUnixMillis(mil int64) time.Time {
	return time.Unix(0, mil*int64(time.Millisecond))
}

const dateTimeLayout = "2006-01-02 15:04:05.000"

func StringToTime(str string) (time.Time, error) {
	return time.Parse(dateTimeLayout, str)
}
