package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func strup(s string) string {
	var res string
	prev := ""
	for pos, char := range s {
		intd, err := strconv.Atoi(string(char))
		if _, errr := strconv.Atoi(string(prev)); errr == nil && err == nil {
			return "not valid"
		}
		if err == nil {

			for i := 1; i < intd; i++ {
				res = res + string(prev)
			}

		} else {
			res = res + string(char)

		}
		if pos > pos-1 {
			prev = string(char)
		}

	}

	return res

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	text = strup(text)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)

	} else {
		fmt.Println(text)
		os.Exit(1)
	}
}
