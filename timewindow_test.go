package timewindow

import (
	"log"
	"testing"
	"time"
)

func Test_timeWindow(t *testing.T) {
	type args struct {
		t     time.Time
		open  time.Time
		close time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			args: args{},
			want: true,
			name: "empty",
		},
		{
			args: args{
				t:     time.Now(),
				open:  time.Now().Add(time.Minute * 5),
				close: time.Now().Add(time.Minute * -5),
			},
			want: false,
			name: "Window is closed and spans midnight",
		},
		{
			args: args{
				t:     time.Now(),
				open:  time.Now().Add(time.Minute * -3),
				close: time.Now().Add(time.Minute * -5),
			},
			want: true,
			name: "Window is open and spans midnight",
		},
		{
			args: args{
				t:     time.Now(),
				open:  time.Now().Add(time.Minute * -1),
				close: time.Now().Add(time.Minute * 2),
			},
			want: true,
			name: "Window is open",
		},
		{
			args: args{
				t:     time.Now(),
				open:  time.Now(),
				close: time.Now(),
			},
			want: true,
			name: "Same Entries",
		},
		{
			args: args{
				t:     time.Now(),
				open:  time.Now().Add(time.Minute * 1),
				close: time.Now().Add(time.Minute * 3),
			},
			want: false,
			name: "Window opens 1 minute from now",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.name, tt.args.t.Unix(), tt.args.open.Unix(), tt.args.close.Unix())
			got := WindowOpen(tt.args.t, tt.args.open, tt.args.close)
			if got != tt.want {
				t.Errorf("windowOpen = %v, want %v", got, tt.want)
			}
		})
	}
}
