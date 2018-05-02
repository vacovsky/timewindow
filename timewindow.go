package timewindow

import (
	"time"
)

// WindowOpen returns a boolean value indicating whether the t value is
// inside of the open window.  If the open time comes after the close time,
// then the window spans midnight (example: window is open from 10PM until 2AM).
// However, if the open time come before the close time (2AM to 10PM), then the
// behaves naturally.
// A return value of True means the submitted time is inside the open window.
func WindowOpen(t, open, close time.Time) bool {

	now := time.Now()
	uNow := now.Unix()

	// times matched or were nil - bail out
	if open.Unix() == close.Unix() {
		return true
	}

	t1 := time.Date(now.Year(), now.Month(), now.Day(),
		open.Hour(), open.Minute(),
		0, 0, time.Local).Unix()

	t2 := time.Date(now.Year(), now.Month(), now.Day(),
		close.Hour(), close.Minute(),
		0, 0, time.Local).Unix()

	if t2 >= t1 {
		// case: window open and close are same day
		return (uNow >= t1 && uNow <= t2)
	} else if t1 >= t2 {
		// case: window spans midnight, window is open before t1 and after t2
		return (uNow >= t1 || uNow <= t2)
	}

	return false
}
