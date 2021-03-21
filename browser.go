package main

import (
    "flag"
    "log"
    "os/exec"
    "os"
    "fmt"
    "runtime"
)

func openbrowser(url string) {
    var err error

    switch runtime.GOOS {
    case "linux":
        fmt.Printf("xdg-open\n")
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
    programName := os.Args[0]
    fmt.Printf("Programname: %s\n",programName)
    argLength := len(os.Args[1:])
    fmt.Printf("argLength: %d\n", argLength)

    for i, a := range os.Args[1:] {
        fmt.Printf("Arg %d is %s\n", i+1, a) 
    }

    var url string
    flag.StringVar(&url, "url", "", "Specify url, no default value")
    flag.Parse()
    if url != "" {
        fmt.Printf("url: %s\n",url)
        openbrowser(url)
    } else {
        fmt.Printf("Usage: --url=<http....>\n")
    }

    fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
    if runtime.GOOS == "windows" {
        fmt.Println("You are running on Windows")
    } else {
        fmt.Println("You are running on an OS other than Windows")
    }

    os.Exit(0)
}
