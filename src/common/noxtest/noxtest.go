package noxtest

import (
	"testing"

	"nox/v1/common/datapath"
)

func DataPath(t testing.TB) string {
	path := datapath.Path()
	if path == "" {
		t.Skip("cannot detect Nox path and NOX_DATA is not set")
	}
	return path
}
