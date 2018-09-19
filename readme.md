```
----------
PLAIN
----------

Running 10s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    12.23ms   22.44ms 264.66ms   91.16%
    Req/Sec     5.41k     3.15k   22.87k    80.82%
  Latency Distribution
     50%    5.90ms
     75%   10.89ms
     90%   30.37ms
     99%  115.96ms
  648155 requests in 10.08s, 72.94MB read
Requests/sec:  64278.53
Transfer/sec:      7.23MB

----------
LOCKED
----------

Running 10s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.05ms   11.22ms 154.29ms   90.05%
    Req/Sec     4.98k     2.33k   22.57k    79.65%
  Latency Distribution
     50%    6.26ms
     75%   11.15ms
     90%   20.19ms
     99%   57.21ms
  594688 requests in 10.09s, 66.92MB read
Requests/sec:  58946.80
Transfer/sec:      6.63MB

----------
LOCKED - counter
----------

Running 10s test @ http://localhost:8080/cnt
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    17.87ms    3.17ms  55.75ms   84.38%
    Req/Sec     1.85k   146.12     3.33k    80.08%
  Latency Distribution
     50%   17.23ms
     75%   18.79ms
     90%   21.54ms
     99%   28.08ms
  221123 requests in 10.04s, 24.88MB read
Requests/sec:  22015.79
Transfer/sec:      2.48MB

----------
LOCKED - goroutines
----------

Running 10s test @ http://localhost:8080/goroutines
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.64ms    8.79ms 116.77ms   73.92%
    Req/Sec     3.96k     0.91k   13.63k    75.31%
  Latency Distribution
     50%    8.28ms
     75%   13.10ms
     90%   19.89ms
     99%   42.03ms
  475381 requests in 10.09s, 53.50MB read
Requests/sec:  47094.43
Transfer/sec:      5.30MB

----------
QUEUE
----------

Running 10s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     8.90ms   10.14ms 116.55ms   90.10%
    Req/Sec     4.73k     1.61k   14.49k    79.75%
  Latency Distribution
     50%    6.67ms
     75%   11.41ms
     90%   18.92ms
     99%   53.29ms
  567789 requests in 10.09s, 63.90MB read
Requests/sec:  56252.82
Transfer/sec:      6.33MB
```