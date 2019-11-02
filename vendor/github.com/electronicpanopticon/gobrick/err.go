package gobrick

import "log"

/// CheckErr takes an error and an optional list of things, and if the error exists logs the error and the string values
/// for the things and then exits the program. Not sure if this is a good thing or not.
func CheckErr(e error, v ...interface{}) {
	if e != nil {
		log.Fatal(e, v)
	}
}