duration=${1:-10}

baseURL="https://squid-app-m6sgv.ondigitalocean.app"

# go-wrk -c 20 \
#  -d "$duration" -body @./testing/assets/card.json -M POST "$baseURL/echo"

# sleep 60

# go-wrk -c 20 \
#  -d "$duration" -body @./testing/assets/card.bin -M POST "$baseURL/echo"

# sleep 60

go-wrk -c 30 \
 -d "$duration" -body @./testing/assets/card.json -M POST "$baseURL/modify-and-return-json"

sleep 10

go-wrk -c 30 \
 -d "$duration" -body @./testing/assets/card.bin -M POST "$baseURL/modify-and-return-protobuf"
