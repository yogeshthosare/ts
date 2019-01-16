package main

import (
    "ts/listeners"
    "ts/storage"

    //"time"
    "fmt"
)

func main() {

	// create job channel
	maxqueuesize := 5
	maxworkers := 3
	

	db:=storage.InitDbGetConn()
	user_data_chan_queue := make(chan storage.UserData, maxqueuesize)
	
	// create workers
	for i := 1; i <= maxworkers; i++ {


		go func(i int) {
				for user_data := range user_data_chan_queue {
					storage.SaveToDb(db, user_data)
				}
		}(i)
	}

	go listeners.ListeningPortOne(user_data_chan_queue)

    go listeners.ListeningPortTwo()
    var input string
    fmt.Scanln(&input)
}
