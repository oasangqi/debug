package debug

import "testing"

func TestDbug(t *testing.T) {
	want := "Hello, world!"
	if got := Dbug(want); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
