package util

import "log"

// Debug Function for helping logging.
func Debug(s string, m string) {
	log.Printf("%s: %s", s, m)
}

// Fatal logs fatal
func Fatal(e error) {
	log.Fatal(e)
}
