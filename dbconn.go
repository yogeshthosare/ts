package ts

import(
    "ts/strlog"
    "os"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDbGetConn(user string, pwd string, ip string, dbport string, dbname string) *gorm.DB {
    connString := user+":"+pwd+"@tcp("+ip+":"+dbport+")/"+dbname
    //db, err := gorm.Open("mysql", "root:yogesh10@tcp(127.0.0.1:3306)/tsdb?charset=utf8&parseTime=True")
    db, err := gorm.Open("mysql", connString)
    //defer db.Close()
    if err!=nil{
        strlog.CommonLogger.Error("Database Initialization failed:", err.Error())
        os.Exit(1)
    }else{
        strlog.CommonLogger.Info("Database Initialised successfully..")
    }
    
    //create table user_data (id varchar(255), name varchar(255), age varchar(255));
    //db.DropTableIfExists(&UserData{})
    //db.CreateTable(&storage.UserData{})
    return db
}