package models

type Cards struct {
	Cards []Card `json:"cards"`
}

type Card struct {
	Artist            string         `json:"artist"`
	ArtistIds         []string       `json:"artistIds"`
	Availability      []string       `json:"availability"`
	BoosterTypes      []string       `json:"boosterTypes"`
	BorderColor       string         `json:"borderColor"`
	ColorIdentity     []string       `json:"colorIdentity"`
	Colors            []string       `json:"colors"`
	ConvertedManaCost int            `json:"convertedManaCost"`
	EdhrecRank        int            `json:"edhrecRank"`
	EdhrecSaltiness   float64        `json:"edhrecSaltiness"`
	Finishes          []string       `json:"finishes"`
	ForeignData       []ForeignData  `json:"foreignData"`
	FrameVersion      string         `json:"frameVersion"`
	HasFoil           bool           `json:"hasFoil"`
	HasNonFoil        bool           `json:"hasNonFoil"`
	Identifiers       Identifiers    `json:"identifiers"`
	IsReprint         bool           `json:"isReprint"`
	Keywords          []string       `json:"keywords"`
	Language          string         `json:"language"`
	Layout            string         `json:"layout"`
	Legalities        Legalities     `json:"legalities"`
	ManaCost          string         `json:"manaCost"`
	ManaValue         int            `json:"manaValue"`
	Name              string         `json:"name"`
	Number            string         `json:"number"`
	OriginalText      string         `json:"originalText"`
	OriginalType      string         `json:"originalType"`
	Power             string         `json:"power"`
	Printings         []string       `json:"printings"`
	PurchaseUrls      PurchaseUrls   `json:"purchaseUrls"`
	Rarity            string         `json:"rarity"`
	SetCode           string         `json:"setCode"`
	SourceProducts    SourceProducts `json:"sourceProducts"`
	Subtypes          []string       `json:"subtypes"`
	Supertypes        []any          `json:"supertypes"`
	Text              string         `json:"text"`
	Toughness         string         `json:"toughness"`
	Type              string         `json:"type"`
	Types             []string       `json:"types"`
	UUID              string         `json:"uuid"`
	Variations        []string       `json:"variations"`
}
type ForeignData struct {
	FlavorText   string `json:"flavorText"`
	Language     string `json:"language"`
	MultiverseID int    `json:"multiverseId"`
	Name         string `json:"name"`
	Text         string `json:"text"`
	Type         string `json:"type"`
}
type Identifiers struct {
	CardKingdomID          string `json:"cardKingdomId"`
	McmID                  string `json:"mcmId"`
	McmMetaID              string `json:"mcmMetaId"`
	MtgjsonFoilVersionID   string `json:"mtgjsonFoilVersionId"`
	MtgjsonV4ID            string `json:"mtgjsonV4Id"`
	MtgoFoilID             string `json:"mtgoFoilId"`
	MtgoID                 string `json:"mtgoId"`
	MultiverseID           string `json:"multiverseId"`
	ScryfallID             string `json:"scryfallId"`
	ScryfallIllustrationID string `json:"scryfallIllustrationId"`
	ScryfallOracleID       string `json:"scryfallOracleId"`
	TcgplayerProductID     string `json:"tcgplayerProductId"`
}
type Legalities struct {
	Commander       string `json:"commander"`
	Duel            string `json:"duel"`
	Legacy          string `json:"legacy"`
	Modern          string `json:"modern"`
	Oathbreaker     string `json:"oathbreaker"`
	Paupercommander string `json:"paupercommander"`
	Penny           string `json:"penny"`
	Predh           string `json:"predh"`
	Premodern       string `json:"premodern"`
	Vintage         string `json:"vintage"`
}
type PurchaseUrls struct {
	CardKingdom string `json:"cardKingdom"`
	Cardmarket  string `json:"cardmarket"`
	Tcgplayer   string `json:"tcgplayer"`
}
type SourceProducts struct {
	Nonfoil []string `json:"nonfoil"`
}
