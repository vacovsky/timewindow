package timewindow

import (
	"time"
)

type config struct {
	Open  time.Time
	Close time.Time
}

func readConfig() {

}

func windowOpen(conf config) bool {

	now := time.Now()

	t1 := time.Date(now.Year(), now.Month(), now.Day(),
		conf.Open.Hour(), conf.Open.Minute(),
		0, 0, time.Local).Unix()

	t2 := time.Date(now.Year(), now.Month(), now.Day(),
		conf.Close.Hour(), conf.Close.Minute(),
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
