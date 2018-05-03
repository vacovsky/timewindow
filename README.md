# timewindow

A little module to do time comparisons for detecting if a Go time struct is between two other times of day.

## the thing

BetweenTimes(t, open, close) returns a boolean value indicating whether the `t` value is inside of the hours and minutes of the `open` window and the `close`.
 
If the open time comes after the close time, then the window spans midnight (example: window is open from 10PM until 2AM). However, if the open time come before the close time (2AM to 10PM), then it behaves intuitively.

A return value of `true` means the submitted time is inside the window.
