package journal_ring

import (
	"journal-ring/priority"
	"os"
	"testing"
	"time"
)

func TestLogNil(t *testing.T) {
	if journal, err := Open(os.Args[0]); err != nil {
		t.Fatal(err)
	} else {
		defer journal.Close()
		for i := 1; i <= 1000; i++ {
			journal.l(priority.INFO, time.Now(), "journal_ring_test")
		}
	}
}
