package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "os/exec"
    "strings"
    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/start", Start)
    router.HandleFunc("/stop", Stop)
    router.HandleFunc("/export_delay/{delay}",exp_delay)
    router.HandleFunc("/export_expid/{delay}",exp_expid)
    router.HandleFunc("/export_runid/{delay}",exp_runid)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func printCommand(cmd *exec.Cmd) {
  fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
  if err != nil {
    os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
  }
}

func printOutput(outs []byte) {
  if len(outs) > 0 {
    fmt.Printf("==> Output: %s\n", string(outs))
  }
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func Start(w http.ResponseWriter, r *http.Request) {

// Create an *exec.Cmd
cmd := exec.Command("/etc/init.d/thesis","start")

// Combine stdout and stderr
printCommand(cmd)
output, err := cmd.CombinedOutput()
printError(err)
printOutput(output)
        fmt.Fprintln(w,"Successfully started thesis daemon")
}

func Stop(w http.ResponseWriter, r *http.Request) {
    
// Create an *exec.Cmd
cmd := exec.Command("/etc/init.d/thesis","stop")

// Combine stdout and stderr
printCommand(cmd)
output, err := cmd.CombinedOutput()
printError(err)
printOutput(output)
        fmt.Fprintln(w,"Successfully stopped thesis daemon")
}

func exp_delay(w http.ResponseWriter, r *http.Request) {
     vars := mux.Vars(r)
     delay := vars["delay"]   
    // Open a new file for writing only
    file, err := os.OpenFile(
        "delay.txt",
        os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
        0666,
    )
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Write bytes to file
    byteSlice := []byte(delay)
    bytesWritten, err := file.Write(byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Wrote %d bytes.\n", bytesWritten)

     fmt.Fprintln(w,delay)
}

func exp_expid(w http.ResponseWriter, r *http.Request) {
     vars := mux.Vars(r)
     delay := vars["delay"]
    // Open a new file for writing only
    file, err := os.OpenFile(
        "expid.txt",
        os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
        0666,
    )
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Write bytes to file
    byteSlice := []byte(delay)
    bytesWritten, err := file.Write(byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Wrote %d bytes.\n", bytesWritten)

     fmt.Fprintln(w,delay)
}

func exp_runid(w http.ResponseWriter, r *http.Request) {
     vars := mux.Vars(r)
     delay := vars["delay"]
    // Open a new file for writing only
    file, err := os.OpenFile(
        "runid.txt",
        os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
        0666,
    )
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Write bytes to file
    byteSlice := []byte(delay)
    bytesWritten, err := file.Write(byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Wrote %d bytes.\n", bytesWritten)

     fmt.Fprintln(w,delay)
}
