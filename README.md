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
➜  protobufvsjson git:(main) ✗ sh testing/load_test.sh 120
Running 120s test @ http://localhost:8080/echo
  20 goroutine(s) running concurrently
3653896 requests in 1m54.44008009s, 23.59GB read
Requests/sec:           31928.46
Transfer/sec:           211.07MB
Avg Req Time:           626.4µs
Fastest Request:        23.208µs
Slowest Request:        24.049083ms
Number of Errors:       0
Running 120s test @ http://localhost:8080/echo
  20 goroutine(s) running concurrently
4287887 requests in 1m54.849146592s, 16.86GB read
Requests/sec:           37334.95
Transfer/sec:           150.33MB
Avg Req Time:           535.691µs
Fastest Request:        20.541µs
Slowest Request:        73.654959ms
Number of Errors:       0
Running 120s test @ http://localhost:8080/modify-and-return-json
  20 goroutine(s) running concurrently
2430383 requests in 1m59.12466151s, 12.86GB read
Requests/sec:           20402.01
Transfer/sec:           110.53MB
Avg Req Time:           980.295µs
Fastest Request:        82µs
Slowest Request:        57.05375ms
Number of Errors:       0
Running 120s test @ http://localhost:8080/modify-and-return-protobuf
  20 goroutine(s) running concurrently
3666774 requests in 1m56.905463565s, 14.45GB read
Requests/sec:           31365.29
Transfer/sec:           126.53MB
Avg Req Time:           637.647µs
Fastest Request:        31.416µs
Slowest Request:        24.792458ms
Number of Errors:       0
```


## Findings when running against a Digital Ocean App

Context

1. Server is in new york
1. Lowest Pro Plan
    1. 1gb of ram
    1. vCPU
    1. 1 container