package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(5 * time.Second)
	fmt.Println(5 * time.Microsecond)
	fmt.Println(time.ParseDuration("10m5s"))

	fmt.Println(time.Now())
	fmt.Println(time.Date(2023, time.March, 19, 15, 7, 0, 0, time.Local))
	fmt.Println(time.Parse(time.Kitchen, "11:30AM"))

	fmt.Println(time.Unix(1503673200, 0))

	fmt.Println(time.Now().Add(3 * time.Hour))

	fmt.Println(time.Now().Format(time.DateTime))
	fmt.Println(time.Now().Add(time.Hour).Format("2006/01/02 15:4:5"))

	fmt.Println("waiting 5 seconds.")
	after := time.After(5 * time.Second)
	<-after
	fmt.Println("5 seconds passed.")

	for now := range time.Tick(3 * time.Second) {
		fmt.Printf("now: %s\n", now)
	}

}
