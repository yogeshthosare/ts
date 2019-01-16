package storage

import(
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

func (received_data UserData) SaveInMemory() string {
    user_data_store = append(user_data_store, received_data)
    //fmt.Println(user_data_store)
    return ("message received, data stored in memory")
}

func SaveToDb(db *gorm.DB, user_data UserData) {
        db.Create(&user_data)
}

func ServeData(conn net.Conn, db *gorm.DB){
    if(len(user_data_store)==0){
        db.Find(&user_data_store)
    }
    
    for _, user_data := range user_data_store {
        resp, err := json.Marshal(user_data)
        if err != nil {
            panic(err)
        }    
        conn.Write([]byte(string(resp) + "\n"))
    }
}