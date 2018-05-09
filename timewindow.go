package timewindow

import (
	"time"
)

// BetweenTimes(t, open, close time.Time) returns a boolean value indicating whether the `t` value is
// inside of the hours and minutes of the `open` window and the `close`.  If the open time comes
// after the close time, then the window spans midnight (example: window is open from 10PM until 2AM).
// However, if the open time come before the close time (2AM to 10PM), then the
// behaves naturally.
// A return value of `true` means the submitted `t` is inside the window.
func BetweenTimes(t, open, close time.Time) bool {

	tu := t.Unix()

	// times matched or were nil - bail out
	if open.Unix() == close.Unix() {
		return true
	}

	t1 := time.Date(t.Year(), t.Month(), t.Day(),
		open.Hour(), open.Minute(),
		0, 0, time.Local).Unix()

	t2 := time.Date(t.Year(), t.Month(), t.Day(),
		close.Hour(), close.Minute(),
		0, 0, time.Local).Unix()

	if t2 >= t1 {
		// case: window open and close are same day
		return (tu >= t1 && tu < t2)
	} else if t1 >= t2 {
		// case: window spans midnight, window is open before t1 and after t2
		return (tu > t1 || tu <= t2)
	}

	return false
}
