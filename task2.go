package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func test() {
	for i := 0; i <= 4; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Returning from function")
}

func writeToFile(filename string, data string) {
	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	bytes_count, err := f.WriteString(data)
	check(err)
	fmt.Printf("Wrote %d bytes \n", bytes_count)

	return
}

func readFromFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	check(err)
	fmt.Printf("Contents of file \n%s\n", data)

	return
}

func main() {
	filename := "new_file.txt"
	data := string("Wikipedia is a free online encyclopedia, created and edited by volunteers around the world and hosted by the Wikimedia Foundation.")
	writeToFile(filename, data)
	readFromFile(filename)
}
