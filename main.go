package main

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/griggsca91/protobufvsjson/models"
	"github.com/griggsca91/protobufvsjson/models/protos/protobuf"
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
			err := json.Unmarshal(body, &jsonCard)
			if err != nil {
				return err
			}

			jsonCard.Artist = jsonCard.Artist + "modified"
			jsonCard.EdhrecRank += 1

			bytes, err := json.Marshal(jsonCard)
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

		if err = os.WriteFile("./testing/assets/card.bin", bytes, 0x777); err != nil {
			panic(err)
		}
	},
}

var rootCmd = &cobra.Command{}

func main() {
	rootCmd.AddCommand(server)
	rootCmd.AddCommand(createProtobinary)

	rootCmd.Execute()
}
