package timewindow

import (
	"time"
)

func WindowOpen(t, open, close time.Time) bool {

	now := time.Now()

	t1 := time.Date(now.Year(), now.Month(), now.Day(),
		open.Hour(), open.Minute(),
		0, 0, time.Local).Unix()

	t2 := time.Date(now.Year(), now.Month(), now.Day(),
		close.Hour(), close.Minute(),
		0, 0, time.Local).Unix()

	if t1 >= t2 {
		// case: window spans midnight, window is open before t1 and after t2
		if now.Unix() <= t1 || now.Unix() >= t2 {
			return true
		}
	} else if t2 >= t1 {
		// case: window open and close are same day
		if now.Unix() >= t1 || now.Unix() <= t2 {
			return true
		}
	}
	return false
}
