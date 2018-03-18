// Package env provides a convenient way
// to convert environment variables into Go data.
package env // import "github.com/kr/env"

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"time"
)

// Int returns the value of the named environment variable,
// interpreted as an int (using strconv.Atoi).
// If there is an error parsing the value, it prints a
// diagnostic message to the log and calls os.Exit(1).
// If name isn't in the environment, it returns value.
func Int(name string, value int) int {
	if s := os.Getenv(name); s != "" {
		var err error
		value, err = strconv.Atoi(s)
		if err != nil {
			log.Println(name, err)
			os.Exit(1)
		}
	}
	return value
}

// Duration returns the value of the named environment variable,
// interpreted as a time.Duration (using time.ParseDuration).
// If there is an error parsing the value, it prints a
// diagnostic message to the log and calls os.Exit(1).
// If name isn't in the environment, it returns value.
func Duration(name string, value time.Duration) time.Duration {
	if s := os.Getenv(name); s != "" {
		var err error
		value, err = time.ParseDuration(s)
		if err != nil {
			log.Println(name, err)
			os.Exit(1)
		}
	}
	return value
}

// Time returns the value of the named environment variable,
// interpreted as a time.Time
// (using time.Parse with the given format).
// If there is an error parsing the value, it prints a
// diagnostic message to the log and calls os.Exit(1).
// If name isn't in the environment, Time returns the time.Time
// that results from parsing the given value.
// Time panics if there is an error parsing the given value.
func Time(name, format, value string) time.Time {
	v, err := time.Parse(format, value)
	if err != nil {
		panic(err)
	}
	if s := os.Getenv(name); s != "" {
		v, err = time.Parse(format, s)
		if err != nil {
			log.Println(name, err)
			os.Exit(1)
		}
	}
	return v
}

// URL returns the value of the named environment variable,
// interpreted as a *url.URL (using url.Parse).
// If there is an error parsing the environment value, it prints a
// diagnostic message to the log and calls os.Exit(1).
// If name isn't in the environment, URL returns the *url.URL
// that results from parsing the given value.
// URL panics if there is an error parsing the given value.
func URL(name string, value string) *url.URL {
	v, err := url.Parse(value)
	if err != nil {
		panic(err)
	}
	if s := os.Getenv(name); s != "" {
		v, err = url.Parse(s)
		if err != nil {
			log.Println(name, err)
			os.Exit(1)
		}
	}
	return v
}

// String returns the value of the named environment variable.
// If name isn't in the environment or is empty, it returns value.
func String(name string, value string) string {
	if s := os.Getenv(name); s != "" {
		value = s
	}
	return value
}
