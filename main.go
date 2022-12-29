package main

import (
	"fmt"
	"os"
)

func main() {
	//var ntptime , err string
	err := "ss"
	//ntptime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err == "nil" {
		fmt.Println("ss")
		os.Exit(1)
		//fmt.Println(os.Stderr, ntptime)

	} else {
		fmt.Println(err)
		os.Exit(0)
	}
}
