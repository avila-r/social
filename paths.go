package application

import (
	"path/filepath"
	"runtime"
)

var (
	// Retrieves the caller's information (file, line, etc.) at runtime.
	_, b, _, _ = runtime.Caller(0)

	// RootPath represents the application's root directory path.
	RootPath = filepath.Join(
		filepath.Dir(b), "/",
	)

	// E2ePath represents the path to the end-to-end (E2E) tests directory.
	E2ePath = filepath.Join(
		filepath.Dir(b), "/e2e",
	)
)
