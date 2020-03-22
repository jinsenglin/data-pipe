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
    instance = flag.String("instance", "", "The Cloud Spanner Instance within your project.")
    database = flag.String("database", "", "The database name in your Cloud Spanner Instance.")

    numAccounts = flag.Int("accounts", 10000, "Number of accounts to generate / load.")
    numSingers = flag.Int("singers", 10000, "Number of singers to generate / load.")
    numAlbums = flag.Int("albums", 100000, "Number of albums to generate / load.")
    numSongs = flag.Int("songs", 1000000, "Number of songs to generate / load.")
    recordsPerFile = 100000

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
            load()
            break
        case "reset":
            // TODO: Implement.
            break
        default:
            log.Fatalf("'%v' is not a valid command! Supported commands are: create generate load reset", command)
    }
}
