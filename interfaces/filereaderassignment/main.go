package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

	// my version
	fileContents, err := ioutil.ReadFile(myFileName)
	fmt.Println("JONO VERSION=============")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("++++File Contents++++")
	fmt.Println(string(fileContents))

	//instructor's version
	myFile, err := os.Open(myFileName)
	fmt.Println("INSTRUCTOR VERSION=============")
	//os.Open() returns *File pointer
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//File type implements Read function which has a byte slice input https://golang.org/pkg/os/#File
	// func (f *File) Read(b []byte) (n int, err error)

	//io.Copy uses destination of Writer type and source of Reader type (ie... the file)
	//os.Stdout -- Stdin, Stdout, and Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.
	//Open results in myfile (type *File pointer).File pointers type uses Read ([]byte)
	//io.Copy uses Reader type. Reader type also uses Read ([]byte)
	io.Copy(os.Stdout, myFile)

	//other student 1
	fmt.Println("STUDENT VERSION=============")
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	buffer := make([]byte, 32*1024)

	for {
		slice, err := f.Read(buffer)

		if slice > 0 {
			fmt.Print(string(buffer[:slice]))
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("read %d bytes: %v", slice, err)
			break
		}
	}
}
