package ts

import(
    "ts/storage"
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDbGetConn(user string, pwd string, ip string, dbport string, dbname string) *gorm.DB {
    connString := user+":"+pwd+"@tcp("+ip+":"+dbport+")/"+dbname
    //db, err := gorm.Open("mysql", "root:yogesh10@tcp(127.0.0.1:3306)/tsdb?charset=utf8&parseTime=True")
    db, err := gorm.Open("mysql", connString)
    //defer db.Close()
    if err!=nil{
        fmt.Println("Connection Failed to Open")
    }
    fmt.Println("Connection Established")
    //db.DropTableIfExists(&UserData{})
    db.CreateTable(&storage.UserData{})
    return db
}