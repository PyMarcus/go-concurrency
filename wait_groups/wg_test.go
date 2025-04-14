package waitgroups

import (
	"sync"
	"testing"
)

func TestShow(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage("ola", &wg)
	wg.Wait()
	msg := show()
	if msg != "ola" {
		t.Errorf("Test has been failed. Expected: ola ,received: %s", msg)
	}
}
