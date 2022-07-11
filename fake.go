package goose

var fakeApply = false

// SetFakeApply set flag not to run migrations on database
func SetFakeApply(fake bool) {
	fakeApply = fake
}
