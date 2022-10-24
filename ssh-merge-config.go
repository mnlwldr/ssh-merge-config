package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	DefaultInput := fmt.Sprintf("%s/%s", homedir, ".ssh/config.d")
	DefaultOutput := fmt.Sprintf("%s/%s", homedir, ".ssh/config")

	in := flag.String("i", DefaultInput, "input directory")
	out := flag.String("o", DefaultOutput, "output file")
	verbose := flag.Bool("v", false, "verbose")

	flag.Parse()

	directory := *in
	output := *out

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Fatalf("%s does not exist.", directory)
	}

	if *verbose {
		fmt.Printf("Read files from %s\n", *in)
	}
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	var content bytes.Buffer

	for _, f := range files {
		readFile, err := ioutil.ReadFile(directory + string(os.PathSeparator) + f.Name())

		if *verbose {
			log.Printf("Read file %s", f.Name())
		}

		if err != nil {
			log.Fatal(err)
		}

		content.WriteString(string(readFile))
	}

	if _, err := os.Stat(output); err == nil {
		if err := backup(output); err != nil {
			log.Fatal(err)
		}
	}

	outputFile, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(outputFile)

	_, err2 := outputFile.WriteString(content.String())
	if err2 != nil {
		log.Fatal(err2)
	}

	if *verbose {
		fmt.Println("done")
	}
}

func backup(filename string) error {
	var backupFile bytes.Buffer

	backupFile.WriteString(filename)
	backupFile.WriteString(".")
	backupFile.WriteString(time.Now().Format("20060102"))

	err := os.Rename(filename, backupFile.String())
	if err != nil {
		return err
	}
	return nil
}
