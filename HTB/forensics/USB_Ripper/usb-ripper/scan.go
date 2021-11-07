package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("auth.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logg, err := os.Open("syslog")
	if err != nil {
		log.Fatal(err)
	}
	defer logg.Close()

	var f = make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f[scanner.Text()] = scanner.Text()
		// fmt.Println(f)
	}

	lscan := bufio.NewScanner(logg)
	for lscan.Scan() {
		mfg := strings.Split(lscan.Text(), "SerialNumber: ") // Not manufacturer
		if len(mfg) > 1 {
			if _, ok := f[mfg[1]]; !ok {
				fmt.Println(lscan.Text())
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
