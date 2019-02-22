package excercises

import "runtime"

// Version returns the current Go version
func Version() string {
	return runtime.Version()
}
