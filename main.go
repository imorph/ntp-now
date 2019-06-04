package main

import (
	"fmt"
	myntp "github.com/imorph/ntp-now/pkg/myntp"
)

func main() {
	fmt.Println("Current time is: ", myntp.Now())
}