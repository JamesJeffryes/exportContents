package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	secretsPath := *flag.String("f", "test", "Path to folder of files to export")
	flag.Parse()
	log.Printf("Exporting file contents from %s", secretsPath)
	cmd := flag.Args()
	exportContents(secretsPath)
	// Exec the command
	if len(cmd) > 0 {
		err := exec.Command(cmd[0], cmd...).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func exportContents(secretsPath string) {
	fileSlice, err := ioutil.ReadDir(secretsPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fileSlice {

		content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", secretsPath, file.Name()))
		if err != nil {
			log.Println(err)
		}
		log.Printf("Exporting %s", file.Name())
		err = os.Setenv(file.Name(), string(content))
		if err != nil {
			log.Println(err)
		}
	}
}
