package main

import (
    "net"
    "fmt"
    "log"
    "strconv"
    "math/rand"
    "time"
    "flag"
)

var c chan string = make(chan string)

func clientSock(delay time.Duration, req_per_user int){
    var cmd string
    var id string
    timeout := (4 * time.Second)
    for i := 0; i < req_per_user; i++ {
      conn, err := net.DialTimeout("tcp", "localhost:2525", timeout)
      if(err!=nil){
        log.Println(err.Error())
      }
      id = strconv.Itoa(rand.Intn(100000))
	    cmd = "{\"Id\": \""+id+"\", \"Name\": \"John Hasa\", \"Age\": \"27\"}" + "\n"
      conn.Write([]byte(cmd))
      time.Sleep(delay * time.Second)
      conn.Close()
    }
}

func main() {
    var (
      concurrent_user = flag.Int("concurrent-user", 1000, "max concurrent requests, default is 1000")
      req_per_user    = flag.Int("req-per-user", 1000, "requests per user, default is 1000")
      delay           = flag.Duration("delay", 5, "delay between two consecative requests made by a user, put 1 ns for 1 second and likewise")
    )
    
    flag.Parse()

    log.Println(*req_per_user)
    log.Println(*concurrent_user)
    log.Println(*delay)
    
    

    for i := 1; i <= *concurrent_user; i++ {
        go func(i int) {
                clientSock(*delay, *req_per_user)
        }(i)
    }
    var input string
    fmt.Scanln(&input)
}