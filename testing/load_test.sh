duration=${1:-10}

go-wrk -c 20 \
    -d "$duration" -body @./testing/assets/card.json -M POST http://localhost:8080/echo
go-wrk -c 20 \
 -d "$duration" -body @./testing/assets/card.bin -M POST http://localhost:8080/echo

go-wrk -c 20 \
 -d "$duration" -body @./testing/assets/card.json -M POST http://localhost:8080/modify-and-return-json

go-wrk -c 20 \
 -d "$duration" -body @./testing/assets/card.bin -M POST http://localhost:8080/modify-and-return-protobuf
