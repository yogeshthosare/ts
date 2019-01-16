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
