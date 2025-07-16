// cmd/epochnexus/main.go
package main

import (
"flag"
"log"
"os"

"epochnexus/internal/epochnexus"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := epochnexus.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
