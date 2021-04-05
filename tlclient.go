package main

import (
	"flag"
	"fmt"
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

func main() {
	var url string

	flag.StringVar(&url, "u", "", "Specify which url")
	flag.Parse()
	if url == "" {
		fmt.Println("Need -u <url>")
		os.Exit(1)
	}
	go openbrowser(url)
	time.Sleep(1 * time.Second)
	os.Exit(0)
}
