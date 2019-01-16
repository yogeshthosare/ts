package listeners

import (
    "ts"
    "ts/handlers"

    "fmt"
    "net"
    "os"
)

func ListeningPortTwo() {
    //Listen on specified port.
    listener, err := net.Listen(ts.Conn_Protocol, ts.Conn_Host+":"+ts.Conn_Port2)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    //Close listener when application stops.
    defer listener.Close()
    fmt.Println("Listening on " + ts.Conn_Host + ":" + ts.Conn_Port2)

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
        go handlers.HandleRequestTwo(conn)
    }
}
