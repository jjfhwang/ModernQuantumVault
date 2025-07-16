// cmd/modernquantumvault/main.go
package main

import (
"flag"
"log"
"os"

"modernquantumvault/internal/modernquantumvault"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := modernquantumvault.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
