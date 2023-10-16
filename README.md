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

1. Server is in new york
1. Lowest Pro Plan
    1. 1gb of ram
    1. vCPU
    1. 1 container