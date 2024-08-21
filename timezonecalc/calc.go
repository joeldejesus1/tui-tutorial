// The purpose of this package is to calculate different timezones.
// This code takes a given timezone and apply it to the base timezone, which is UTC. Then the code will return the correct time at the given timezone overstdout.
package timezonecalc

import "fmt"

// write a function to print the user arguments
func PrintCLIArgs(args []string) {
	gmtdiff, err := ReadTimeZone(args[0])
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("GMT Diff is %d\n", gmtdiff)
	}
}

// github issue #1
// The code needs to read the input, which will be the timezone the user chooses.
// The output will be the current time at the chosen timezone, using UTC time the base.
func ReadTimeZone(timezone string) (gmtdiff int, err error) {
	// write a map from string timezone to gmtdiff
	timezoneM := make(map[string]int)
	timezoneM["JST"] = -9
	timezoneM["EST"] = -5
	// look up the timezone
	var present bool
	gmtdiff, present = timezoneM[timezone]
	// do an if test
	if present {
		// this part means what?
		err = fmt.Errorf("%s is not a timezone") // put this here?
	} else {
		// this part means what?
	}
	// this line returns gmtdiff and err at the same time
	return
}
