package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func encodeMap() {
	writer, err := os.OpenFile("index.txt", os.O_CREATE|os.O_WRONLY, 0666) // create file if not exists
	check(err)

	empData, err3 := json.Marshal(m)
	if err != nil {
		fmt.Println(err3.Error())
		return
	}

	_, err2 := writer.Write(empData)
	check(err2)
	defer writer.Close()

}

func decodeMap(done chan bool) {
	data, err3 := ioutil.ReadFile("index.txt")
	check(err3)

	if err2 := json.Unmarshal(data, &m); err2 != nil {
		panic(err2)
	}
	done <- true
}
