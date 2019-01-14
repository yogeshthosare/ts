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

