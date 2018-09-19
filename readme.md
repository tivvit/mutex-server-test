```
----------

PLAIN
----------

Running 10s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.58ms   22.02ms 464.78ms   93.22%
    Req/Sec     4.81k     2.35k   21.33k    81.22%
  Latency Distribution
     50%    6.61ms
     75%   11.52ms
     90%   23.22ms
     99%  105.19ms
  577009 requests in 10.10s, 64.93MB read
Requests/sec:  57143.60
Transfer/sec:      6.43MB

----------
LOCKED
----------

Running 10s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.24ms   10.46ms 119.67ms   89.90%
    Req/Sec     4.55k     1.74k   25.22k    84.70%
  Latency Distribution
     50%    6.85ms
     75%   11.58ms
     90%   19.80ms
     99%   54.14ms
  544608 requests in 10.10s, 61.29MB read
Requests/sec:  53922.69
Transfer/sec:      6.07MB

----------
LOCKED - counter
----------

Running 10s test @ http://localhost:8080/cnt
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    18.14ms    3.63ms  52.71ms   83.59%
    Req/Sec     1.83k   158.08     4.47k    79.67%
  Latency Distribution
     50%   17.48ms
     75%   19.36ms
     90%   22.46ms
     99%   29.64ms
  218600 requests in 10.06s, 24.60MB read
Requests/sec:  21721.23
Transfer/sec:      2.44MB

----------
LOCKED - goroutines
----------

Running 10s test @ http://localhost:8080/goroutines
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    10.25ms    9.14ms 229.16ms   70.30%
    Req/Sec     3.63k   763.64     7.31k    73.33%
  Latency Distribution
     50%    9.08ms
     75%   14.39ms
     90%   21.50ms
     99%   41.21ms
  435527 requests in 10.08s, 49.01MB read
Requests/sec:  43214.68
Transfer/sec:      4.86MB
```
