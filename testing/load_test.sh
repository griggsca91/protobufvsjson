duration=${1:-10}

baseURL="http://24.144.82.58:8080"

echo "\n\nEchoing json\n\n"
go-wrk -c 10 \
 -d "$duration" -body @./testing/assets/card.json -M POST "$baseURL/echo"

sleep 10

echo "\n\nEchoing string heavy protobuf\n\n"
go-wrk -c 10 \
 -d "$duration" -body @./testing/assets/card.bin -M POST "$baseURL/echo"

sleep 10

echo "\n\nEchoing enum and strings protobuf\n\n"
go-wrk -c 10 \
 -d "$duration" -body  @./testing/assets/card_with_enum.bin -M POST "$baseURL/echo"

sleep 10
echo "\n\nModify and return json\n\n"


go-wrk -c 10 \
 -d "$duration" -body @./testing/assets/card.json -M POST "$baseURL/modify-and-return-json"

sleep 10
echo "\n\nModify and return string protobuf\n\n"


go-wrk -c 10 \
 -d "$duration" -body @./testing/assets/card.bin -M POST "$baseURL/modify-and-return-protobuf"


sleep 10

echo "\n\nModify and return enumed protobuf\n\n"

go-wrk -c 10 \
 -d "$duration" -body @./testing/assets/card_with_enum.bin -M POST "$baseURL/modify-and-return-protobuf-with-enums"
