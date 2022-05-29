package goose

var unsafeApply = false

// AllowUnsafeMigrations set flag to apply unsafe migrations
func AllowUnsafeMigrations(unsafe bool) {
	unsafeApply = unsafe
}
