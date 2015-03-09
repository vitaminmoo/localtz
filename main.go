package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type format struct {
	name   string
	layout string
}

func (f format) String() string {
	return fmt.Sprintf("%s=%q", f.name, f.layout)
}

var formats = []format{
	{"LogStd", "2006/01/02 15:04:05"},
	{"LogMicro", "2006/01/02 15:04:05.000000"},
	{"ANSIC", "Mon Jan _2 15:04:05 2006"},
	{"UnixDate", "Mon Jan _2 15:04:05 MST 2006"},
	{"RubyDate", "Mon Jan 02 15:04:05 -0700 2006"},
	{"RFC822", "02 Jan 06 15:04 MST"},
	{"RFC822Z", "02 Jan 06 15:04 -0700"}, // RFC822 with numeric zone
	{"RFC850", "Monday, 02-Jan-06 15:04:05 MST"},
	{"RFC1123", "Mon, 02 Jan 2006 15:04:05 MST"},
	{"RFC1123Z", "Mon, 02 Jan 2006 15:04:05 -0700"}, // RFC1123 with numeric zone
	{"RFC3339", "2006-01-02T15:04:05Z07:00"},
	{"RFC3339Nano", "2006-01-02T15:04:05.999999999Z07:00"},
	{"Kitchen", "3:04PM"},
	// "Handy" time stamps.
	{"Stamp", "Jan _2 15:04:05"},
	{"StampMilli", "Jan _2 15:04:05.000"},
	{"StampMicro", "Jan _2 15:04:05.000000"},
	{"StampNano", "Jan _2 15:04:05.000000000"},
}

func main() {
	log.SetFlags(0)
	loc, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatalf("Unable to load local timezone: %v", err)
	}
	locflag := loc.String()
	flag.StringVar(&locflag, "l", locflag, "location")

	verbose := false
	flag.BoolVar(&verbose, "v", verbose, "toggle verbose output")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Println("missing time argument")
		flag.PrintDefaults()
	}

	if locflag != loc.String() {
		loc, err = time.LoadLocation(locflag)
		if err != nil {
			log.Fatalf("Unable to load timezone %s: %v", locflag, err)
		}
	}

	for _, tf := range formats {
		t, err := time.Parse(tf.layout, args[0])
		if err == nil {
			if verbose {
				log.Printf("Parsed using %s", tf)
			}
			fmt.Printf("From %s: %s     To %s: %s\n", t.Location(), args[0], loc, t.In(loc).Format(tf.layout))
			return
		}
		if verbose {
			log.Printf("Tried %s: %v", tf, err)
		}

	}
	os.Exit(1)
}
