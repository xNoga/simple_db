package main

import (
	"os"
	"strconv"
)

var m = make(map[int]string)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	done := make(chan bool, 1)
	go decodeMap(done)
	<-done

	if os.Args[1] == "set" {
		done = make(chan bool, 1)
		writeLine(os.Args[2], done)
		<-done
		encodeMap()
	}

	if os.Args[1] == "get" {
		index, err := strconv.ParseInt(os.Args[2], 10, 64)
		check(err)

		offset, err2 := strconv.ParseInt(m[int(index)], 10, 64)
		check(err2)
		readLine(offset)
	}
}
