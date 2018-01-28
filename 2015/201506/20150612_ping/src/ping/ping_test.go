package ping

import (
	"testing"
)

func Test_ping(ts *testing.T) {
	ret := Pinger("127.0.0.1", 10)
	if ret == nil {
		ts.Logf("OK\n")
	} else {
		ts.Logf("failed: %v\n", ret)
	}
}
