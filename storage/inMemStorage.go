package storage

import(
	"fmt"
    "log"
    "net"
    "encoding/json"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserData struct {
    Id string `json:"id"`
    Name string `json: "name"`
    Age string `json: age`
}

var user_data_store []UserData

func InitDbGetConn() *gorm.DB {
    db, err := gorm.Open("mysql", "root:yogesh10@tcp(127.0.0.1:3306)/tsdb?charset=utf8&parseTime=True")
    //defer db.Close()
    if err!=nil{
        fmt.Println("Connection Failed to Open")
        log.Println(err)
    }
    fmt.Println("Connection Established")
    //db.DropTableIfExists(&UserData{})
    db.CreateTable(&UserData{})
    return db
}

func (received_data UserData) SaveInMemory() string {
    user_data_store = append(user_data_store, received_data)
    //fmt.Println(user_data_store)
    return ("message received, data stored in memory")
}

func SaveToDb(db *gorm.DB, user_data UserData) {
           db.Create(&user_data)
}

func ServeData(conn net.Conn){
    //var user_data UserData
    //if(UserData{}==user_data){}
    for _, user_data := range user_data_store {
        //db.Create(&user_data_store)
        resp, err := json.Marshal(user_data)
        if err != nil {
            panic(err)
        }    
        conn.Write([]byte(string(resp) + "\n"))
    }
}