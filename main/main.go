package main

import (
    "ts/listeners"
    "fmt"
)

func main() {
    go listeners.ListeningPortOne()
    //go listeningPortTwo()
    var input string
    fmt.Scanln(&input)
}
