package timing_test

import (
	"github.com/nikitych1w/learn-go-patterns/Profiling/timing/timing"
	"testing"
	"time"
)

func TestTimeProfiler_Duration(t *testing.T) {
	defer timing.Duration(time.Now(), "first")
	time.Sleep(time.Second * 5)
}
