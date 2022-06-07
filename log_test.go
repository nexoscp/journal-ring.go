package journal_ring

import (
	"journal-ring/priority"
	"os"
	"testing"
)

func TestLogNil(t *testing.T) {
	if journal, err := Open(os.Args[0]); err != nil {
		t.Fatal(err)
	} else {
		journal.l(priority.INFO, "journal_ring_test")
	}
}
