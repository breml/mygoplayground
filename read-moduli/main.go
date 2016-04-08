package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("/etc/ssh/moduli")
	if err != nil {
		fmt.Errorf("Error while reading /etc/ssh/moduli\n")
	}

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		moduli := strings.Split(line, " ")
		if len(moduli) == 7 {
			fmt.Printf("{ %s, %s, %s, mustBigIntFromHexString(\"%s\") },\n", moduli[0], moduli[4], moduli[5], moduli[6])
		}
	}

}
