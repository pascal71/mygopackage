// Tests for mygopackage
package mygopackage

import (
	"testing"
)

// Test the Greet function
func TestGreet(t *testing.T) {
	name := "Alice"
	got := Greet(name)
	want := "Hello, Alice, greetings from mygopackage v" + version

	if got != want {
		t.Errorf("Greet(%q) = %q; want %q", name, got, want)
	}
}
