package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//os.Args returns slice of args including the args for running golang itself
	//[C:\Users\JONO-W~1\AppData\Local\Temp\go-build083238294\b001\exe\main.exe instructions.txt]
	//the file name is then the second item in the array
	myFileName := os.Args[1]

	//OpenFile is the generalized open call; most users will use Open or Create instead. https://golang.org/pkg/os/#Open
	// myFile, err := os.OpenFile(myFileName, os.O_RDWR|os.O_CREATE, 0755)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if err := myFile.Close(); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(os.O_RDWR)
	// fmt.Println(os.O_CREATE)

	fileContents, err := ioutil.ReadFile(myFileName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("++++File Contents++++")
		fmt.Println(string(fileContents))
	}
}
