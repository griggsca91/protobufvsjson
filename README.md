# protobufvsjson

## Introduction

Simple comparison of sending and receiving JSON and Protobuf.  Goal was to do a simple test to see if there was a clear and distinct winner in performance for sending data as well as decoding and encoding.

## To Run Locally

In one terminal 
```sh
go run . server
```

In another terminal
```sh
sh testing/load_test.sh 20
```

## To Do

Launch this up on a DO server and run go-wrk on it 


## Current Findings when running locally
While not a good test where the load testing environment is the same as the environment being tested, it's at least a consistent environment to compare against each other

```sh
➜  protobufvsjson git:(main) ✗ sh testing/load_test.sh 1


Echoing json


Running 1s test @ http://localhost:8080/echo
  20 goroutine(s) running concurrently
29745 requests in 960.377336ms, 187.22MB read
Requests/sec:           30972.20
Transfer/sec:           194.95MB
Avg Req Time:           645.74µs
Fastest Request:        23.125µs
Slowest Request:        8.970917ms
Number of Errors:       0


Echoing string heavy protobuf


Running 1s test @ http://localhost:8080/echo
  20 goroutine(s) running concurrently
36153 requests in 965.62765ms, 145.57MB read
Requests/sec:           37439.90
Transfer/sec:           150.75MB
Avg Req Time:           534.189µs
Fastest Request:        24.625µs
Slowest Request:        13.1765ms
Number of Errors:       0


Echoing enum and strings protobuf


Running 1s test @ http://localhost:8080/echo
  20 goroutine(s) running concurrently
49297 requests in 962.761992ms, 189.84MB read
Requests/sec:           51203.72
Transfer/sec:           197.18MB
Avg Req Time:           390.596µs
Fastest Request:        20.25µs
Slowest Request:        12.588583ms
Number of Errors:       0


Modify and return json


Running 1s test @ http://localhost:8080/modify-and-return-json
  30 goroutine(s) running concurrently
19809 requests in 990.542286ms, 107.32MB read
Requests/sec:           19998.14
Transfer/sec:           108.35MB
Avg Req Time:           1.500139ms
Fastest Request:        89.375µs
Slowest Request:        33.713458ms
Number of Errors:       0


Modify and return string protobuf


Running 1s test @ http://localhost:8080/modify-and-return-protobuf
  30 goroutine(s) running concurrently
29071 requests in 973.124967ms, 117.27MB read
Requests/sec:           29873.86
Transfer/sec:           120.51MB
Avg Req Time:           1.004222ms
Fastest Request:        34.625µs
Slowest Request:        18.553333ms
Number of Errors:       0


Modify and return enumed protobuf


Running 1s test @ http://localhost:8080/modify-and-return-protobuf-with-enums
  30 goroutine(s) running concurrently
36171 requests in 971.998499ms, 139.57MB read
Requests/sec:           37213.02
Transfer/sec:           143.59MB
Avg Req Time:           806.169µs
Fastest Request:        34.75µs
Slowest Request:        13.522917ms
Number of Errors:       0
```


## Findings when running against a Digital Ocean App

Context

1. Two Servers both located in San Francisco
1. Low end stuff, nothing special


```sh
root@tester:~/protobufvsjson# sh testing/load_test.sh 60


Echoing json


Running 60s test @ http://10.124.0.2:8080/echo
  20 goroutine(s) running concurrently
182788 requests in 59.070836903s, 1.12GB read
Requests/sec:           3094.39
Transfer/sec:           19.48MB
Avg Req Time:           6.463316ms
Fastest Request:        550.269µs
Slowest Request:        127.365183ms
Number of Errors:       0


Echoing string heavy protobuf


Running 60s test @ http://10.124.0.2:8080/echo
  20 goroutine(s) running concurrently
254381 requests in 59.217658583s, 1.00GB read
Requests/sec:           4295.70
Transfer/sec:           17.30MB
Avg Req Time:           4.655824ms
Fastest Request:        454.051µs
Slowest Request:        48.876096ms
Number of Errors:       0


Echoing enum and strings protobuf


Running 60s test @ http://10.124.0.2:8080/echo
  20 goroutine(s) running concurrently
351464 requests in 59.269298639s, 1.32GB read
Requests/sec:           5929.95
Transfer/sec:           22.84MB
Avg Req Time:           3.372709ms
Fastest Request:        410.7µs
Slowest Request:        55.119054ms
Number of Errors:       0


Modify and return json


Running 60s test @ http://10.124.0.2:8080/modify-and-return-json
  20 goroutine(s) running concurrently
208461 requests in 59.421336484s, 1.10GB read
Requests/sec:           3508.18
Transfer/sec:           19.01MB
Avg Req Time:           5.700954ms
Fastest Request:        590.751µs
Slowest Request:        67.144438ms
Number of Errors:       0


Modify and return string protobuf


Running 60s test @ http://10.124.0.2:8080/modify-and-return-protobuf
  20 goroutine(s) running concurrently
279121 requests in 59.303596663s, 1.10GB read
Requests/sec:           4706.65
Transfer/sec:           18.99MB
Avg Req Time:           4.24931ms
Fastest Request:        479.463µs
Slowest Request:        42.834719ms
Number of Errors:       0


Modify and return enumed protobuf


Running 60s test @ http://10.124.0.2:8080/modify-and-return-protobuf-with-enums
  20 goroutine(s) running concurrently
360810 requests in 59.2383683s, 1.36GB read
Requests/sec:           6090.82
Transfer/sec:           23.50MB
Avg Req Time:           3.283632ms
Fastest Request:        399.415µs
Slowest Request:        211.267484ms
```