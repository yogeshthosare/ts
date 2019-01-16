package listeners

import (
    "ts/handlers"
    
    "fmt"
    "net"
    "os"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func ListeningPortTwo(hostip string , port2 string ,db *gorm.DB) {
    //Listen on specified port.
    listener, err := net.Listen("tcp",  hostip+":"+port2)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    //Close listener when application stops.
    defer listener.Close()
    fmt.Println("Listening on " + hostip + ":" + port2)

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
        go handlers.HandleRequestTwo(conn, db)
    }
}
