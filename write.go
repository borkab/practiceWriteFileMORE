package main

import (
	"bufio"
	"fmt"
	"os"
)

//just practice how to write data to a file
//according to https://gobyexample.com/writing-files

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//how to dump a string (or just bytes) into a file
	d1 := []byte("Good\tmorning\tsunshine")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	//you can open or create a file for writing
	hifile, err := os.Create("plopp.txt")
	check(err)

	defer hifile.Close()

	//you can write by slices
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := hifile.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	//WriteString()
	n3, err := hifile.WriteString("dancing in the moonlight\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	//you can sync to delete writes to stable storage
	hifile.Sync()

	//bufio provides buffered writers too
	w := bufio.NewWriter(hifile)
	n4, err := w.WriteString("take me home\ncountry road")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	//use FLush() to ensure all buffered operations have ben applied to the underlying writer
	w.Flush()
}
