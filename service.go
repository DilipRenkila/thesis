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
//    router.HandleFunc("/exp_delay/{delay}",exp_delay)
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


