A] Problem statement
Service exposes a simple tcp server that listens on two ports

i. First Port: (Default port 2525)
Accepts a JSON payload Example: {“id”: “1”, “name”: “John Doe”, “age”: “25”}
and saves it to a buffered storage, say an array in memory. Another background
process asynchronously writes this data to a database or a File.

ii. Second Port: (Default port 2526)
Any request on another Port Reads from the buffered storage, if the data exists,
the endpoints send the response back. If the data does not exist in the buffered
storage, it reads the data from the database/file and dumps the output and prints
records line-by-line.

B] Approach - 

I have maintained a queuesize and workers, data will be fetched into the queue and then
safely saved into database by worker in background asynchronously. 

As stated in problem statement data is also stored in memory. Any request on second port 
fetches the data from in-memory if avalable else fetches it from database and dumps output
line by line.

Service is designed in structured way depending on the functionality of each component. 

I have used gorm library for better db performance.

Log file is maintained seperately.

C] How to run main service ?

Please make sure to create database and table using provided sql script.

	create database tsdb;
	use tsdb;
	create table user_data (id varchar(255), name varchar(255), age varchar(255));

create binary of your service 

	go build -o main main/main.go

	./main --help

	Usage of ./main:

 	-dbip string
    	Database ip (default "127.0.0.1")
 	-dbname string
    	Database name (default "tsdb")
 	-dbport string
    	Database port (default "3306")
 	-dbpwd string
    	Database password (default "yogesh10")
 	-dbuser string
    	Database username (default "root")
 	-logLevel string
    	possible levels -debug, info, warn|warning, error, fatal, panic (default "info")
 	-port1 string
    	Port accepting json and storing (default "2525")
 	-port2 string
    	Port reading from storage and serving (default "2526")
 	-queue-size int
    	Max records in a channel queue (default 1000)
 	-server-ip string
    	Host ip (default "localhost")
 	-worker-size int
    	Max goroutins for database entries (default 10)


D] How to run test script ?

use shell script provided in test_scripts directory directly

	./basic_test.sh
	
	for i in {1..1000} ; do
  	echo {\"Id\": \"$i\", \"Name\": \"John Hasa\", \"Age\": \"30\"} | nc localhost 2525 &
	done
	
Any request on second port fetches the data from in-memory if avalable else fetches it from database and dumps output
line by line.
	
  	echo "get data" | nc localhost 2526

use golang script for concurent testing and performance testing

	go build -o [test_script] [test_scripts/test_client.go]

	./test_client --help

Usage of test_client:

  	-concurrent-user int
    max concurrent requests, default is 1000 (default 1000)
			
  	-delay duration
    delay between two consecative requests made by a user (default 5ns)
			
  	-req-per-user int
		requests per user, default is 1000 (default 1000)

E] Performance and observation


./test_client -concurrent-user 1 -delay 0ns -req-per-user 100000 -> 15k requests per minute

./test_client -concurrent-user 250 -delay 1ns -req-per-user 1000 -> 14k requests per minute

./test_client -concurrent-user 500 -delay 2ns -req-per-user 1000 -> 16k requests per minute

./test_client -concurrent-user 1000 -delay 3ns -req-per-user 1000 -> 25k requests per minute

Have to put delay to wait for opened connections to be closed.
Please let me know your feedback. Thanks you it was indeed a great problem to spend time on.
