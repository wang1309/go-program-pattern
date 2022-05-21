package func_options

import (
	"testing"
	"time"
)

func TestOption(t *testing.T) {
	_, _ = NewServer("localhost", 1024)
	_, _ = NewServer("localhost", 2048, Protocol("udp"))
	_, _ = NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))
}
