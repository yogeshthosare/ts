package main

import (
    "ts/listeners"
    "ts/storage"
    "ts"
    "fmt"
    "sync"
    "flag"
)

func main() {

	var (
		maxqueuesize = flag.Int("queue-size", 100, "Max records in a channel queue")
		maxworkers   = flag.Int("worker-size", 5, "Max goroutins for database entries")
		port1        = flag.String("port1", "2525", "Port accepting json and storing")
		port2        = flag.String("port2", "2526", "Port reading from storage and serving")
		hostip	 	 = flag.String("server-ip", "localhost", "Host ip")

		//database
		user         = flag.String("dbuser", "root", "Database username")
		pwd          = flag.String("dbpwd", "yogesh10", "Database password")
		dbip         = flag.String("dbip", "127.0.0.1", "Database ip")
		dbport       = flag.String("dbport", "3306", "Database port")
		dbname       = flag.String("dbname", "tsdb", "Database name")
	)

	flag.Parse()

	var wg sync.WaitGroup

	db:=ts.InitDbGetConn(*user, *pwd, *dbip, *dbport, *dbname)
	user_data_chan_queue := make(chan storage.UserData, *maxqueuesize)

	// create workers
	wg.Add(*maxworkers)
	for i := 1; i <= *maxworkers; i++ {
		wg.Add(1)
		go func(i int) {
				for user_data := range user_data_chan_queue {
					storage.SaveToDb(db, user_data)
				}
		}(i)
	}

	go listeners.ListeningPortOne(*hostip, *port1, user_data_chan_queue)

    go listeners.ListeningPortTwo(*hostip, *port2,  db)
    var input string
    fmt.Scanln(&input)
}
