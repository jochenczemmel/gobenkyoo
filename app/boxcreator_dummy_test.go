package app_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/app"
)

// TestCreateboxLoadStore uses a test dummy to check error handling.
func TestCreateboxLoadStore(t *testing.T) {
	testCases := []struct {
		name            string
		loadStorer      app.ClassroomLoadStorer
		wantOK          bool
		wantLoadErrMsg  string
		wantStoreErrMsg string
	}{{
		name:       "ok",
		loadStorer: dummy{},
		wantOK:     true,
	}, {
		name:            "ClassroomLoadStorer is nil",
		wantLoadErrMsg:  "no ClassroomLoadStorer defined",
		wantStoreErrMsg: "no ClassroomLoadStorer defined",
	}, {
		name:           "load lib returns error",
		loadStorer:     dummy{loadError: "load failed"},
		wantLoadErrMsg: "load failed",
	}, {
		name:           "load room returns error",
		loadStorer:     dummy{loadRoomError: "load room failed"},
		wantLoadErrMsg: "load room failed",
	}, {
		name:           "load returns path error",
		loadStorer:     dummy{pathError: "file does not exist"},
		wantLoadErrMsg: "open .: file does not exist",
	}, {
		name:       "load room returns path error",
		loadStorer: dummy{roomPathError: "file does not exist"},
	}, {
		name:            "store error",
		loadStorer:      dummy{storeRoomError: "store room failed"},
		wantOK:          true,
		wantStoreErrMsg: "store room failed",
	}}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			creator := app.NewBoxCreator(c.loadStorer)

			t.Run("load", func(t *testing.T) {
				ok, err := creator.Load("", "")
				checkErrorMessage(t, err, c.wantLoadErrMsg)
				if ok != c.wantOK {
					t.Errorf("ERROR: got %v, want %v", ok, c.wantOK)
				}
			})

			t.Run("store", func(t *testing.T) {
				err := creator.Store()
				checkErrorMessage(t, err, c.wantStoreErrMsg)
			})
		})
	}
}
