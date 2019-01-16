for i in {500..1500} ; do
  echo {\"Id\": \"$i\", \"Name\": \"John Hasa\", \"Age\": \"30\"} | nc localhost 2525
done
