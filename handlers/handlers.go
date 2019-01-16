package handlers

import (
    "ts/strlog"
    "ts/storage"
    "net"
    "os"
    //"log"
    "encoding/json"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)


//Handler for listener one
func HandleRequestOne(conn net.Conn, user_data_chan_queue chan storage.UserData) {
  // Buffer to hold incoming stream of data.
  buffer := make([]byte, 1024)
  //required data types
  var user_data storage.UserData

  // Read the incoming connection into the buffer.
  lenReq, err := conn.Read(buffer)
  
  if err != nil {
    strlog.CommonLogger.Error("Error reading on first handler: ", err.Error())
    os.Exit(1)
  }
  conn.Write(buffer)

  //get only actual received data
  received_data := (buffer[:lenReq])

  if err := json.Unmarshal(received_data, &user_data); err != nil {
      panic(err)
  }

  // putting data in channel 
  user_data_chan_queue <- user_data

  user_data.SaveInMemory()
  
  //Send response back to the caller
  conn.Write([]byte("data received and stored\n"))
  
  //Close the connection after execution.
  conn.Close()
}


//Handler for listener one
func HandleRequestTwo(conn net.Conn, db *gorm.DB) {
  // Buffer to hold incoming stream of data.
  buffer := make([]byte, 1024)
  
  // Read the incoming connection into the buffer.
  lenReq, err := conn.Read(buffer)
  if err != nil {
    strlog.CommonLogger.Error("Error reading on second handler: ", err.Error())
    os.Exit(1)
  }

  if lenReq>0{
    //Send response back to the caller
    //conn.Write([]byte("any request on port two received\n"))
    storage.ServeData(conn, db)
  }
  //Close the connection after execution.
  conn.Close()
}