package internal

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	RootPath = filepath.Join(
		filepath.Dir(b), "../",
	)

	E2ePath = filepath.Join(
		filepath.Dir(b), "../e2e",
	)
)
