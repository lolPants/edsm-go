package edsm

import "testing"

func BoolToQueryTest(t *testing.T) {
	got := boolToQuery(true)
	if got != "1" {
		t.Errorf("boolToQuery(true) = %s", got)
	}

	got = boolToQuery(false)
	if got != "0" {
		t.Errorf("boolToQuery(true) = %s", got)
	}
}

func BoolToQueryBenchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boolToQuery(i%2 == 0)
	}
}
