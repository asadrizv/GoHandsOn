package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("ma_first_file.txt")
	if err != nil {

	}
	file.WriteString("hahahahahah")
	file.Close()

	file2, err := os.Create("ma_second_file.txt")
	if err != nil {

	}
	file2.WriteString("hahahahahah")
	file2.Close()

	checkfile1, err := ioutil.ReadFile("ma_first_file.txt")
	checkfile2, err := ioutil.ReadFile("ma_second_file.txt")

	h := crc32.NewIEEE()
	h2 := crc32.NewIEEE()

	h.Write(checkfile1)
	h2.Write(checkfile2)

	fmt.Println(h.Sum32() == h2.Sum32())

}
