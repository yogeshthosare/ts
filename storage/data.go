package storage

import(
    "ts/strlog"
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

func (received_data UserData) SaveInMemory() {
    user_data_store = append(user_data_store, received_data)
    strlog.CommonLogger.Info("Data saved in-memory..", received_data)
}

func SaveToDb(db *gorm.DB, user_data UserData) {
    db.Create(&user_data)
    strlog.CommonLogger.Info("Data saved into database..", user_data)
}

func ServeData(conn net.Conn, db *gorm.DB){
    if(len(user_data_store)==0){
        strlog.CommonLogger.Info("Data served from database ..")
        db.Find(&user_data_store)
        if(len(user_data_store)==0){
            strlog.CommonLogger.Info("Data served from database ..")
            conn.Write([]byte("No data available in storage! " + "\n"))
        }
    }else{
        strlog.CommonLogger.Info("Data served from in-memory ..")
    }
    
    for _, user_data := range user_data_store {
        resp, err := json.Marshal(user_data)
        if err != nil {
            panic(err)
        }    
        conn.Write([]byte(string(resp) + "\n"))
    }
}