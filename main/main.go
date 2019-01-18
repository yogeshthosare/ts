package main

import (
	"ts"
    "ts/listeners"
    "ts/storage"
    "ts/strlog"
    "fmt"
    "flag"
    "os"
  	log "github.com/sirupsen/logrus"
)

func init() {
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&log.JSONFormatter{})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  log.SetLevel(log.WarnLevel)

  
}



func main() {

	var (
		maxqueuesize = flag.Int("queue-size", 1000, "Max records in a channel queue")
		maxworkers   = flag.Int("worker-size", 10, "Max goroutins for database entries")
		port1        = flag.String("port1", "2525", "Port accepting json and storing")
		port2        = flag.String("port2", "2526", "Port reading from storage and serving")
		hostip	 	 = flag.String("server-ip", "localhost", "Host ip")

		//database
		user         = flag.String("dbuser", "root", "Database username")
		pwd          = flag.String("dbpwd", "yogesh10", "Database password")
		dbip         = flag.String("dbip", "127.0.0.1", "Database ip")
		dbport       = flag.String("dbport", "3306", "Database port")
		dbname       = flag.String("dbname", "tsdb", "Database name")

		//log level
		reqLevel     = flag.String("logLevel", "info", "possible levels -debug, info, warn|warning, error, fatal, panic")
	)

	flag.Parse()

	//by default log level
	log.SetLevel(log.InfoLevel)

	//set log level
	logLevel, err := log.ParseLevel(*reqLevel)
	if err != nil{
  		strlog.CommonLogger.Info("Unable to set LogLevel, setting 'Info' by default")
	}else{
  		log.SetLevel(logLevel)
	}

	file, err := os.OpenFile("log/logger.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
  	if err != nil {
        log.Fatal(err)
  	}
  	//set output of log to log file
  	defer file.Close()
  	log.SetOutput(file)

	strlog.CommonLogger.Info("Initializing database.. ")
	db:=ts.InitDbGetConn(*user, *pwd, *dbip, *dbport, *dbname)


	user_data_chan_queue := make(chan storage.UserData, *maxqueuesize)

	// create workers
	//Background process asynchronously writes this data to a database or a File.
	for i := 1; i <= *maxworkers; i++ {
		go func(i int) {
				for user_data := range user_data_chan_queue {
					storage.SaveToDb(db, user_data)
				}
		}(i)
	}

	//listen on port one
	//Accepts a JSON payload and saves it to a buffered storage, say an array in memory.
	go listeners.ListeningPortOne(*hostip, *port1, user_data_chan_queue)

	//listen on port two
	//Any request on another Port Reads from the buffered storage
    go listeners.ListeningPortTwo(*hostip, *port2,  db)

    //saving process from stopping 
    var input string
    fmt.Scanln(&input)
}
