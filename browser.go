package main

import (
    "io/ioutil"
    "io"
    "flag"
    "log"
    "os/exec"
    "os"
    "fmt"
    "runtime"
    "path"
)


//    "fmt"
//    "io/ioutil"
//    "os"

// File copies a single file from src to dst
func File(src, dst string) error {
    var err error
    var srcfd *os.File
    var dstfd *os.File
    var srcinfo os.FileInfo

    if srcfd, err = os.Open(src); err != nil {
        return err
    }
    defer srcfd.Close()

    if dstfd, err = os.Create(dst); err != nil {
        return err
    }
    defer dstfd.Close()

    if _, err = io.Copy(dstfd, srcfd); err != nil {
        return err
    }
    if srcinfo, err = os.Stat(src); err != nil {
        return err
    }
    return os.Chmod(dst, srcinfo.Mode())
}

//Copy a directory recursively

// Dir copies a whole directory recursively
func Dir(src string, dst string) error {
    var err error
    var fds []os.FileInfo
    var srcinfo os.FileInfo

    if srcinfo, err = os.Stat(src); err != nil {
        return err
    }

    if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
        return err
    }

    if fds, err = ioutil.ReadDir(src); err != nil {
        return err
    }
    for _, fd := range fds {
        srcfp := path.Join(src, fd.Name())
        dstfp := path.Join(dst, fd.Name())

        if fd.IsDir() {
            if err = Dir(srcfp, dstfp); err != nil {
                fmt.Println(err)
            }
        } else {
            if err = File(srcfp, dstfp); err != nil {
                fmt.Println(err)
            }
        }
    }
    return nil
}

func readdir(dir string) {
   // Open the directory.
    outputDirRead, _ := os.Open(dir)

    // Call Readdir to get all files.
    outputDirFiles, _ := outputDirRead.Readdir(0)

    // Loop over files.
    for outputIndex := range(outputDirFiles) {
        outputFileHere := outputDirFiles[outputIndex]

        // Get name of file.
        outputNameHere := outputFileHere.Name()

        // Print name.
        fmt.Println(outputNameHere)
    }
}

func copy(src, dst string) (int) {
    input, err := ioutil.ReadFile(src)
    if err != nil {
        fmt.Println(err)
        return 1
    }

    os.Remove(dst)

    err = ioutil.WriteFile(dst, input, 0644)
    if err != nil {
        fmt.Println("Error creating", dst)
        fmt.Println(err)
        return 1
    }
    return 0
}

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

    hostname, hosterr := os.Hostname()
    if hosterr != nil {
        fmt.Println(hosterr)
        os.Exit(1)
    }
    fmt.Printf("Hostname: %s\n", hostname)

    for i, a := range os.Args[1:] {
        fmt.Printf("Arg %d is %s\n", i+1, a) 
    }

    readdir(".")

    copy("apa","bepa")
    err := os.RemoveAll("/tmp/bepa")
    if err != nil { 
        log.Fatal(err) 
    } 
    //os.Exit(2)
    Dir("/tmp/apa","/tmp/bepa")

    var url string
    flag.StringVar(&url, "url", "", "Specify url, no default value")
    flag.Parse()
    if url != "" {
        fmt.Printf("Opening url: %s\n",url)
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
