package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	val  string
	next int
}

func readRange(vals map[int]node, start, end int) string {
	var o string
	i := start
	for i != end {
		o += vals[i].val
		i = vals[i].next
	}
	if i == end {
		o += vals[end].val
	}
	return o
}

func main() {
	file, err := os.Open("deterministic.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var f = make(map[int]node)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		index, _ := strconv.Atoi(vals[0])
		next, _ := strconv.Atoi(vals[2])
		stringy, _ := strconv.Atoi(vals[1])
		stringy = stringy ^ 0x69
		f[index] = node{string(stringy), next}
	}

	out := readRange(f, 69420, 999)
	fmt.Println(out)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
