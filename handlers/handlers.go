package handlers

import (
    "ts"

    "fmt"
    "net"
    "encoding/json"
)

//Handler for listener one
func HandleRequest(conn net.Conn) {
  
  //required data types
  var user_data ts.UserData

  // Buffer to hold incoming stream of data.
  buffer := make([]byte, 1024)
  
  // Read the incoming connection into the buffer.
  lenReq, err := conn.Read(buffer)
  
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }
  conn.Write(buffer)

  //get only actual received data
  received_data := (buffer[:lenReq])

  if err := json.Unmarshal(received_data, &user_data); err != nil {
      panic(err)
  }
  fmt.Println(user_data.Id, user_data.Name, user_data.Age)

  //Send response back to the caller
  conn.Write([]byte("message received\n"))
  //Close the connection after execution.
  conn.Close()
}