Expose a simple TCP (not-HTTP) server that listens on two ports.

Will keep on adding repo details and steps for running the service here.

First commit progress -
a. basic initial structure is ready, may change as I proceed.

b. service is now able to listen on a port, read incoming data and responds back to the caller
   how to run : 
	just run the binary without any parameters, service listens on localhost port 2525
   how to request : from command line
	echo '{"Id": "1", "Name": "John Hasa", "Age": "27"}'  | nc localhost 2525

c. data handling, data manipulating, data storing, error handling, logging other port functioning etc to be done.

Second commit progress - 
a. some changes in structure, may have to work on variable names, strutures etc.. working out functionality. 

b. able to put data in local in memory storage, also persisting data in mysql db

c. creating n number of workers (goroutines) to persist data in db, enqueing incoming data/records in maxqueue (channel)

d. able to serve in-memory data to the caller on port 2526 on any request, as mentioned in problem statement

e. pending tasks - serving data from db if not present in-memory, error handling, logging, accepting parameters from user while running binary, setting accurate max-workers and max-queue size

f. this is not final commit.. will update this info if needed.

Third commit progress

a. serving data from db if not present in-memory, error handling, logging, accepting parameters from user while running binary, setting accurate max-workers and max-queue size

b. some structure changes

c. performance tesing pending 

d. updated way to run the binary -


	Usage of ../../bin/main:
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
    	Max records in a channel queue (default 100)
  -server-ip string
    	Host ip (default "localhost")
  -worker-size int
    	Max goroutins for database entries (default 5)


   eg. $./main -port1 3030 -port 3031 

e. I have added one basic script to test the service

   for i in {1..10000} ; do
  	echo {\"Id\": \"$i\", \"Name\": \"John Hasa\", \"Age\": \"30\"} | nc localhost 2525
   done

   run this shell script of simply feed one by one data


f. golang script to test service with concurrent users making request.
	please check test_scripts.go and use it for testing.
   	how to run - 
	Usage of ../../bin/test_client:
  -concurrent-user int
    	Concurrent requests (default 1000)
  -delay duration
    	Max goroutins for database entries (default 5ns)
  -req-per-user int
    	Max goroutins for database entries (default 1000)



