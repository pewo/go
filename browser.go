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

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
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
	fmt.Println("url is", url)
	go openbrowser(url)
	time.Sleep(1 * time.Second)
	os.Exit(0)
}
