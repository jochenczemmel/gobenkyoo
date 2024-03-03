package csvimport_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/store/csvimport"
)

func TestWordFormat(t *testing.T) {
	testCases := []struct {
		name    string
		input   []string
		wantErr bool
	}{{
		name:  "ok",
		input: []string{"NIHONGO", "", "KANA"},
	}, {
		name:    "wrong content",
		input:   []string{"NIHONGO", "CHINESE"},
		wantErr: true,
	}, {
		name:    "no fields",
		input:   []string{},
		wantErr: true,
	}, {
		name:    "nil fields",
		wantErr: true,
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			_, err := csvimport.NewWordFormat(c.input...)
			if c.wantErr {
				if err == nil {
					t.Fatalf("ERROR: wanted error not detected")
				}
				t.Logf("INFO: got error: %v", err)
				return
			}
			if err != nil {
				t.Fatalf("ERROR: got error: %v", err)
			}
		})
	}
}
