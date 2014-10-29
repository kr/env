// Package env provides a convenient way to convert environment
// variables into Go data. It is similar in design to package
// flag.
package env

import (
	"log"
	"os"
	"strconv"
)

var funcs []func() bool

// Int returns a new int pointer.
// When env.Parse is called,
// env var name will be parsed
// and the resulting value
// will be assigned to the location pointed to.
func Int(name string, value int) *int {
	p := new(int)
	*p = value
	funcs = append(funcs, func() bool {
		if s := os.Getenv(name); s != "" {
			v, err := strconv.Atoi(s)
			if err != nil {
				log.Println(name, err)
				return false
			}
			*p = v
		}
		return true
	})
	return p
}

// String returns a new string pointer.
// When env.Parse is called,
// env var name will be assigned
// to the location pointed to.
func String(name string, value string) *string {
	p := new(string)
	*p = value
	funcs = append(funcs, func() bool {
		if s := os.Getenv(name); s != "" {
			*p = s
		}
		return true
	})
	return p
}

// Parse parses known env vars
// and assigns the values to the variables
// that were previously registered.
func Parse() {
	ok := true
	for _, f := range funcs {
		ok = f() && ok
	}
	if !ok {
		os.Exit(1)
	}
}
