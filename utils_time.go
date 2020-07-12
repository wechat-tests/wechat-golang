package wxxx

import "time"

var location = time.FixedZone("UTC", int((time.Hour * 8).Seconds()))

const longTimeFormat = "2006-01-02 15:04:05"
const longTimeWithMsFormat = "2006-01-02 15:04:05.999999999"

func TimeNow() time.Time {
	now := LocalTime(time.Now())
	return now
}
func LocalTime(t time.Time) time.Time {
	return t.In(location)
}
