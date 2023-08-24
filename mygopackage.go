// Package mygopackage provides simple greeting utilities.
package mygopackage

import "fmt"

const version = "0.1.0"

// Greet returns a greeting message for the provided name.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s, greetings from mygopackage v%s", name, version)
}
