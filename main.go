package main

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/griggsca91/protobufvsjson/models"
	"github.com/griggsca91/protobufvsjson/models/protos/protobuf"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var server = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New()
		app.Get("/metrics", monitor.New())

		app.Post("echo", func(c *fiber.Ctx) error {
			return c.Send(c.Body())
		})

		app.Post("modify-and-return-json", func(c *fiber.Ctx) error {

			body := c.Body()

			var jsonCard models.Card
			if err := json.Unmarshal(body, &jsonCard); err != nil {
				return err
			}

			jsonCard.Artist = jsonCard.Artist + "modified"
			jsonCard.EdhrecRank += 1

			bytes, err := json.Marshal(&jsonCard)
			if err != nil {
				return err
			}

			return c.Send(bytes)
		})

		app.Post("modify-and-return-protobuf", func(c *fiber.Ctx) error {

			body := c.Body()

			var protoCard protobuf.Card
			if err := proto.Unmarshal(body, &protoCard); err != nil {
				return err
			}

			protoCard.Artist = protoCard.Artist + "modified"
			protoCard.EdhrecRank += 1

			bytes, err := proto.Marshal(&protoCard)
			if err != nil {
				return err
			}

			return c.Send(bytes)
		})

		app.Post("modify-and-return-protobuf-with-enums", func(c *fiber.Ctx) error {

			body := c.Body()

			var protoCard protobuf.CardWithEnum
			if err := proto.Unmarshal(body, &protoCard); err != nil {
				return err
			}

			protoCard.Artist = protoCard.Artist + "modified"
			protoCard.EdhrecRank += 1

			bytes, err := proto.Marshal(&protoCard)
			if err != nil {
				return err
			}

			return c.Send(bytes)
		})

		app.Listen(":8080")

	},
}

var createProtobinary = &cobra.Command{
	Use: "gen-proto-bin",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile("./testing/assets/card.json")
		if err != nil {
			panic(err)
		}
		var card protobuf.Card
		if err = json.Unmarshal(data, &card); err != nil {
			panic(err)
		}

		bytes, err := proto.Marshal(&card)
		if err != nil {
			panic(err)
		}

		if err = overwriteFile("./testing/assets/card.bin", bytes); err != nil {
			panic(err)
		}

		cardWithEnum := protobuf.CardWithEnum{
			Artist:    card.Artist,
			ArtistIds: card.ArtistIds,
			Availability: lo.Map(card.Availability, func(availability string, i int) protobuf.CardWithEnum_Availability {
				switch availability {
				case "mgto":
					return protobuf.CardWithEnum_Availability_MTGO
				case "paper":
					return protobuf.CardWithEnum_Availability_Paper
				default:
					return protobuf.CardWithEnum_Availability_Unknown
				}
			}),
			BoosterTypes: lo.Map(card.BoosterTypes, func(boosterType string, _ int) protobuf.CardWithEnum_BoosterType {
				switch boosterType {
				case "default":
					return protobuf.CardWithEnum_BoosterType_Default
				default:
					return protobuf.CardWithEnum_BoosterType_Unknown
				}
			}),
			BorderColor: func(borderColor string) protobuf.CardWithEnum_BorderColor {
				switch borderColor {
				case "black":
					return protobuf.CardWithEnum_BorderColor_Black
				default:
					return protobuf.CardWithEnum_BorderColor_Unknown
				}
			}(card.BorderColor),
			ColorIdentity: lo.Map(card.ColorIdentity, func(colorIdentity string, _ int) protobuf.CardWithEnum_ColorIdentity {
				switch colorIdentity {
				case "white":
					return protobuf.CardWithEnum_ColorIdentify_White
				default:
					return protobuf.CardWithEnum_ColorIdentity_Unknown
				}
			}),
			Colors: lo.Map(card.ColorIdentity, func(colors string, _ int) protobuf.CardWithEnum_Colors {
				switch colors {
				case "white":
					return protobuf.CardWithEnum_Colors_White
				default:
					return protobuf.CardWithEnum_Colors_Unknown
				}
			}),
			ConvertedManaCost: card.ConvertedManaCost,
			EdhrecRank:        card.EdhrecRank,
			EdhrecSaltiness:   card.EdhrecSaltiness,
			Finishes: lo.Map(card.Finishes, func(s string, _ int) protobuf.CardWithEnum_Finishes {
				switch s {
				case "nonfoil":
					return protobuf.CardWithEnum_Finishes_NonFoil
				default:
					return protobuf.CardWithEnum_Finishes_Unknown
				}
			}),
			ForeignData:  card.ForeignData,
			FrameVersion: card.FrameVersion,
			HasFoil:      card.HasFoil,
			HasNonFoil:   card.HasNonFoil,
			Identifiers:  card.Identifiers,
			IsReprint:    card.IsReprint,
			Keywords: lo.Map(card.Keywords, func(s string, _ int) protobuf.CardWithEnum_Keyword {
				switch s {
				case "First strike":
					return protobuf.CardWithEnum_Keyword_FirstStrike
				default:
					return protobuf.CardWithEnum_Keyword_Unknown
				}
			}),
			Language: card.Language,
			Layout: func(s string) protobuf.CardWithEnum_Layout {
				switch s {
				case "normal":
					return protobuf.CardWithEnum_Layout_Normal
				default:
					return protobuf.CardWithEnum_Layout_Unknown

				}
			}(card.Layout),
			Legalities: func(c *protobuf.Card) *protobuf.CardWithEnum_Legalities {

				m := func(s string) protobuf.CardWithEnum_Legalities_Legality {
					switch s {
					case "legal":
						return protobuf.CardWithEnum_Legalities_Legality_Legal
					case "restricted":
						return protobuf.CardWithEnum_Legalities_Legality_Restricted
					default:
						return protobuf.CardWithEnum_Legalities_Legality_Unknown
					}
				}

				l := protobuf.CardWithEnum_Legalities{
					Commander:       m(c.Legalities.Commander),
					Duel:            m(c.Legalities.Duel),
					Legacy:          m(c.Legalities.Legacy),
					Modern:          m(c.Legalities.Modern),
					Oathbreaker:     m(c.Legalities.Oathbreaker),
					Paupercommander: m(c.Legalities.Paupercommander),
					Penny:           m(c.Legalities.Penny),
					Predh:           m(c.Legalities.Predh),
					Premodern:       m(c.Legalities.Premodern),
					Vintage:         m(c.Legalities.Vintage),
				}
				return &l
			}(&card),
			ManaCost:     card.ManaCost,
			ManaValue:    card.ManaValue,
			Name:         card.Name,
			Number:       card.Number,
			OriginalText: card.OriginalText,
			OriginalType: card.OriginalType,
			Power:        card.Power,
			Printings:    card.Printings,
			PurchaseUrls: card.PurchaseUrls,
			Rarity: func(s string) protobuf.CardWithEnum_Rarity {
				switch s {
				case "uncommon":
					return protobuf.CardWithEnum_Rarity_Uncommon
				default:
					return protobuf.CardWithEnum_Rarity_Unknown
				}
			}(card.Rarity),
			SetCode:        card.SetCode,
			SourceProducts: card.SourceProducts,
			Subtypes: lo.Map(card.Subtypes, func(s string, _ int) protobuf.CardWithEnum_SubType {
				switch s {
				case "human":
					return protobuf.CardWithEnum_SubType_Human
				case "cleric":
					return protobuf.CardWithEnum_SubType_Cleric
				default:
					return protobuf.CardWithEnum_SubType_Unknown
				}
			}),
			Supertypes: card.Supertypes,
			Text:       card.Text,
			Toughness:  card.Toughness,
			Type:       card.Type,
			Types: lo.Map(card.Types, func(s string, _ int) protobuf.CardWithEnum_Type {
				switch s {
				case "creature":
					return protobuf.CardWithEnum_Type_Creature
				default:
					return protobuf.CardWithEnum_Type_Unknown
				}
			}),
			Variations: card.Variations,
		}

		bytes, err = proto.Marshal(&cardWithEnum)
		if err != nil {
			panic(err)
		}

		if err = overwriteFile("./testing/assets/card_with_enum.bin", bytes); err != nil {
			panic(err)
		}
	},
}

func overwriteFile(file string, bytes []byte) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	if _, err = f.Write(bytes); err != nil {
		return err
	}
	return nil
}

var rootCmd = &cobra.Command{}

func main() {
	rootCmd.AddCommand(server)
	rootCmd.AddCommand(createProtobinary)

	rootCmd.Execute()
}
