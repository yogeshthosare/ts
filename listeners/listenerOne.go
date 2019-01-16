package listeners

import (
    "ts"
    "ts/handlers"
    "ts/storage" 
    "fmt"
    "net"
    "os"
)

func ListeningPortOne(user_data_chan_queue chan storage.UserData) {
    //Listen on specified port.
    listener, err := net.Listen(ts.Conn_Protocol, ts.Conn_Host+":"+ts.Conn_Port1)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    //Close listener when application stops.
    defer listener.Close()
    fmt.Println("Listening on " + ts.Conn_Host + ":" + ts.Conn_Port1)

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
