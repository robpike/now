// Copyright 2012 The rspace Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Now tells what time it is in other time zones. The first argument
// identifies a time zone either by shorthand (EST, NYC) or by time zone
// file base name, such as Yellowknife or Paris.
//
// % now Paris
// Thu Apr 12 15:55:55 CEST 2012 Paris
// % now Adelaide
// Thu Apr 12 23:26:14 CST 2012 Adelaide
// %

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	zone := ""
	t := time.Now()
	if len(os.Args) > 1 {
		zone = os.Args[1]
		if tz, ok := timeZone[zone]; ok {
			zone = tz
		}
		t = t.In(loadZone(zone))
	}
	fmt.Printf("%s %s\n", t.Format(time.UnixDate), zone)
}

func loadZone(zone string) *time.Location {
	loc, err := time.LoadLocation(zone)
	if err == nil {
		return loc
	}
	// Pure ASCII, but OK. Allow us to say "paris" as well as "Paris".
	if len(zone) > 0 && 'a' <= zone[0] && zone[0] <= 'z' {
		zone = string(zone[0]+'A'-'a') + string(zone[1:])
	}
	// See if there's a file with that name in /usr/share/zoneinfo
	files, _ := filepath.Glob("/usr/share/zoneinfo/*/" + zone)
	if len(files) >= 1 {
		if len(files) > 1 {
			fmt.Fprintf(os.Stderr, "now: multiple time zones; using first of %v\n", files)
		}
		loc, err = time.LoadLocation(files[0][len("/usr/share/zoneinfo/"):])
		if err == nil {
			return loc
		}
	}
	fmt.Fprintf(os.Stderr, "now: %s\n", err)
	os.Exit(1)
	return nil

}

// from /usr/share/zoneinfo
var timeZone = map[string]string{
	"GMT":     "Europe/London",
	"BST":     "Europe/London",
	"BSDT":    "Europe/London",
	"CET":     "Europe/Paris",
	"UTC":     "",
	"PST":     "America/Los_Angeles",
	"PDT":     "America/Los_Angeles",
	"LA":      "America/Los_Angeles",
	"LAX":     "America/Los_Angeles",
	"MST":     "America/Denver",
	"MDT":     "America/Denver",
	"CST":     "America/Chicago",
	"CDT":     "America/Chicago",
	"Chicago": "America/Chicago",
	"EST":     "America/New_York",
	"EDT":     "America/New_York",
	"NYC":     "America/New_York",
	"NY":      "America/New_York",
	"AEST":    "Australia/Sydney",
	"AEDT":    "Australia/Sydney",
	"AWST":    "Australia/Perth",
	"AWDT":    "Australia/Perth",
	"ACST":    "Australia/Adelaide",
	"ACDT":    "Australia/Adelaide",
}
