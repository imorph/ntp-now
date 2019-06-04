package main

import (
	"fmt"
	"os"

	myntp "github.com/imorph/ntp-now/pkg/myntp"
)

func main() {
	ntpTime, err := myntp.Now()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Anyway, here is your LOCAL time: ", ntpTime)
		os.Exit(1)
	}
	
	fmt.Println("Current NETWORK time according to time.apple.com is: ", ntpTime)
}