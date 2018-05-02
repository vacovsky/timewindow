package timewindow

import (
	"testing"
	"time"
)

func Test_timeWindow(t *testing.T) {
	type args struct {
		conf config
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			args: args{
				conf: config{},
			},
			want: true,
			name: "empty",
		},
		{
			args: args{
				conf: config{
					Open:  time.Now().Add(time.Minute * 5),
					Close: time.Now().Add(time.Minute * 1),
				},
			},
			want: false,
			name: "Window is closed and spans midnight",
		},
		{
			args: args{
				conf: config{
					Open:  time.Now().Add(time.Minute * -3),
					Close: time.Now().Add(time.Minute * -5),
				},
			},
			want: true,
			name: "Window is open and spans midnight",
		},
		{
			args: args{
				conf: config{
					Open:  time.Now().Add(time.Minute * -1),
					Close: time.Now().Add(time.Minute * 2),
				},
			},
			want: true,
			name: "Window is open",
		},
		{
			args: args{
				conf: config{
					Open:  time.Now(),
					Close: time.Now(),
				},
			},
			want: true,
			name: "Same Entries",
		},
		{
			args: args{
				conf: config{
					Open:  time.Now().Add(time.Minute * 1),
					Close: time.Now().Add(time.Minute * 3),
				},
			},
			want: false,
			name: "Window opens 1 minute from now",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := windowOpen(tt.args.conf)
			if got != tt.want {
				t.Errorf("windowOpen = %v, want %v", got, tt.want)
			}
		})
	}
}
