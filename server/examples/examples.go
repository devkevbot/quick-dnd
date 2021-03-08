package examples

// This package is intended to be a "playground" for example snippets of
// code.
//
// Code in this package should probably not be used for the actual
// project; it is here simply to illustrate concepts in Go.

// User describes a very simple user model. The `json:...` line is
// called a "tag" and is used when the struct is encoded to JSON or when
// JSON is decoded into this struct.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
