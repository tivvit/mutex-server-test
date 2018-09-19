go build main.go
./main &
siege http://localhost:8080 -c 1000 -t10s