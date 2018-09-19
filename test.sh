echo "----------"
echo "PLAIN"
echo "----------"
echo ""
go build server.go
./server > log &
sleep 1
wrk -t12 -c400 -d10s --latency http://localhost:8080
#tail log
kill %1

go build locked.go
./locked > log &
sleep 1
#siege http://localhost:8080 -c 255 -t10S -r2 -b
echo ""
echo "----------"
echo "LOCKED"
echo "----------"
echo ""
wrk -t12 -c400 -d10s --latency http://localhost:8080
echo ""
echo "----------"
echo "LOCKED - counter"
echo "----------"
echo ""
wrk -t12 -c400 -d10s --latency http://localhost:8080/cnt
echo ""
echo "----------"
echo "LOCKED - goroutines"
echo "----------"
echo ""
wrk -t12 -c400 -d10s --latency http://localhost:8080/goroutines
kill %1

#cat log