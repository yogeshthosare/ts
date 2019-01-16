package listeners

import (
    "ts/handlers"
    "ts/storage" 
    "fmt"
    "net"
    "os"
)

func ListeningPortOne(hostip string, port1 string, user_data_chan_queue chan storage.UserData) {
    //Listen on specified port.
    listener, err := net.Listen("tcp", hostip+":"+port1)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    //Close listener when application stops.
    defer listener.Close()
    fmt.Println("Listening on " + hostip + ":" + port1)

    for {
        //Wait for a connection.
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        //Handle the connection in a new goroutine.
        //The loop then returns to accepting, so that
        //multiple connections may be served concurrently.
        go handlers.HandleRequestOne(conn, user_data_chan_queue)
    }
}
