package timer

import (
	"time"
)

type Timer struct {
	ms_time_out int64
	out         bool
}

// Init initializes its time out ms timestamp
func (t *Timer) Init(MS_timestamp int64) {
	t.ms_time_out = time.Now().UnixMilli() + MS_timestamp
}

// TimeOut checks if it's time out
func (t *Timer) TimeOut() bool {
	if !t.out && time.Now().UnixMilli() > t.ms_time_out {
		t.out = true
	}

	return t.out
}
