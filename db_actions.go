package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func readLine(byteOffset int64) {
	f, err := os.OpenFile("data.txt", os.O_CREATE|os.O_RDONLY, 0666) // create file if not exists
	reader := bufio.NewReader(f)                                     // creates a new reader

	_, err = f.Seek(byteOffset, 0) // Set the current position for the fd
	if err != nil {
		log.Println(err)
		return
	}

	// read line until a new line is found
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(line))
}

func writeLine(input string, done chan bool) {
	fmt.Println("writeLine starting")
	data, err := ioutil.ReadFile("data.txt")
	check(err)

	writer, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) // create file if not exists
	check(err)
	_, err = writer.Write([]byte(input + "\n"))
	check(err)

	if len(m) == 0 {
		m[1] = "0"
	} else {
		fmt.Println(len(data))
		m[len(m)+1] = strconv.Itoa((len(data)))
	}

	defer writer.Close()
	done <- true
}
