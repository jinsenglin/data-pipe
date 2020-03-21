package main

import (
    "log"

    "github.com/namsral/flag"
)

var (
    version string // set by linker -X
    command string // set by first argument

    project = flag.String("project", "", "Your cloud project ID.")
    bucketName = flag.String("bucket", "", "The name of the bucket within your project.")

    debugger debugging
)

func main() {
    // init var
    flag.Parse()
    command = flag.Arg(0)

    // init var
    debugger = debugging(true)

    debugger.DumpVariables()

    switch command {
        case "create":
            // TODO: Implement.
            break
        case "generate":
            generate()
            break
        case "load":
            // TODO: Implement.
            break
        case "reset":
            // TODO: Implement.
            break
        default:
            log.Fatalf("'%v' is not a valid command! Supported commands are: create generate load reset", command)
    }
}
