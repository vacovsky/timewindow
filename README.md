# timewindow

A little module to do time comparisons for detecting if a Go time struct is between two other times of day.

## the thing

 BetweenTimes(time, open, close) returns a boolean value indicating whether the `time` value is
 inside of the hours and minutes of the `open` window and the `close`.  If the open time comes 
 after the close time, then the window spans midnight (example: window is open from 10PM until 2AM).
 However, if the open time come before the close time (2AM to 10PM), then the
 behaves naturally.
 A return value of True means the submitted time is inside the window.
