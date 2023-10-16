package models

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/griggsca91/protobufvsjson/models/protos/protobuf"
	"google.golang.org/protobuf/proto"
)

func TestJSONAndProtoAreSame(t *testing.T) {

	data, err := os.ReadFile("../testing/assets/card.json")
	if err != nil {
		t.Fatal(err)
	}

	var jsonCard Card
	err = json.Unmarshal(data, &jsonCard)
	if err != nil {
		t.Fatal(err)
	}

	data, err = os.ReadFile("../testing/assets/card.bin")
	if err != nil {
		t.Fatal(err)
	}

	var protoCard protobuf.Card
	err = proto.Unmarshal(data, &protoCard)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, jsonCard.Artist, protoCard.Artist)
	assert.Equal(t, jsonCard.Colors, protoCard.Colors)
	assert.Equal(t, jsonCard.Name, protoCard.Name)
	assert.Equal(t, jsonCard.PurchaseUrls.CardKingdom, protoCard.PurchaseUrls.CardKingdom)
	assert.Equal(t, jsonCard.Variations, protoCard.Variations)
	assert.Equal(t, jsonCard.Text, protoCard.Text)

}
