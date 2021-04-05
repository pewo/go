package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func opentlclient(url string) {
	var err error

	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	}
	if err != nil {
		log.Fatal(err)
	}
}

func writeConf() string {
	// Create our Temp File:  This will create a filename like /tmp/prefix-123456
	// We can use a pattern of "pre-*.txt" to get an extension like: /tmp/pre-123456.txt
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix*.tlclient")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}

	// Remember to clean up the file afterwards
	//defer os.Remove(tmpFile.Name())

	fmt.Println("Created File: " + tmpFile.Name())

	// Example writing to the file
	text := []byte("AUTOLOGIN=1\n")
	if _, err = tmpFile.Write(text); err != nil {
		log.Fatal("Failed to write to temporary file", err)
	}

	// Close the file
	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}

	return tmpFile.Name()
}

func main() {
	//var url string

	//flag.StringVar(&url, "u", "", "Specify which url")
	//flag.Parse()
	//if url == "" {
	//	fmt.Println("Need -u <url>")
	//	os.Exit(1)
	//}

	config := writeConf()

	go opentlclient(config)
	time.Sleep(1 * time.Second)
	os.Exit(0)
}
