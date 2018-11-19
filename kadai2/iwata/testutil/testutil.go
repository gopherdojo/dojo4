package testutil

import (
	"strings"
	"testing"
)

func ContainsError(t *testing.T, gotErr, wantErr error, msg string) {
	t.Helper()

	switch {
	case gotErr == nil && wantErr == nil:
		return
	case gotErr == nil:
		t.Fatalf("%s: want [%s] error, but got nil", msg, wantErr)
	case wantErr == nil:
		t.Fatalf("%s: got [%s] error, but want nil", msg, gotErr)
	}

	if !strings.Contains(gotErr.Error(), wantErr.Error()) {
		t.Errorf("%s: [%s] not contains [%s]", msg, gotErr, wantErr)
	}
}
