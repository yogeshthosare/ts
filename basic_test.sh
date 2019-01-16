for i in {1..10000} ; do
  echo {\"Id\": \"$i\", \"Name\": \"John Hasa\", \"Age\": \"30\"} | nc localhost 3030
done
