package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {

	name := os.Args[0]

	if len(os.Args) < 3 {
		help(name)
		os.Exit(255)
	}

	directory := os.Args[1]
	output := os.Args[2]

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Printf("%s does not exist", directory)
		os.Exit(255)
	}

	files, err := ioutil.ReadDir(directory)
	checkErr(err)

	var content bytes.Buffer
	for _, f := range files {
		readFile,err := ioutil.ReadFile(directory + string(os.PathSeparator) + f.Name())
		checkErr(err)
		content.WriteString(string(readFile))
	}

	if _, err := os.Stat(output); err == nil {
		backup(output)
	}

	out, err := os.Create(output)
	checkErr(err)
	defer func(out *os.File) {
		err := out.Close()
		checkErr(err)
	}(out)

	_, err2 := out.WriteString(content.String())
	checkErr(err2)

	fmt.Println("done")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func help(name string) {
	fmt.Printf("Usage: %s [input directory] [output file]\n", name)
	fmt.Println("Example: ")
	fmt.Printf("\t %s ~/.ssh/config.d/ ~/.ssh/config\n", name)
}

func backup(filename string) {
	var backupFile bytes.Buffer
	backupFile.WriteString(filename)
	backupFile.WriteString(".")
	backupFile.WriteString(time.Now().Format("20060102"))

	err := os.Rename(filename, backupFile.String())
	if err != nil {
		checkErr(err)
	}
}
