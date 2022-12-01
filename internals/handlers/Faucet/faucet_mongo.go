package faucetHandler

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/go-playground/validator/v10"
	// "github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/visveshdojima/faucet-backend/database"
	"github.com/visveshdojima/faucet-backend/internals/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

// get faucet data

var faucetCollection *mongo.Collection = database.GetCollection(database.DB_mongo, "faucet")

func GetFaucetFromMongo(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cursor, err := faucetCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Catchphrases Not found",
			"error":   err,
		})
	}

	var blockData []model.Faucet

	for cursor.Next(ctx) {
		var b model.Faucet
		cursor.Decode(&b)
		blockData = append(blockData, b)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Block data Found from mongoDB", "data": blockData})

}

type FaucetResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

var validate = validator.New()

// Creating new faucet item
func CreateFaucetData(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var f model.Faucet
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&f); err != nil {
		return c.Status(http.StatusBadRequest).JSON(FaucetResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// //use the validator library to validate required fields
	// if validationErr := validate.Struct(&f); validationErr != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(FaucetResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	// }

	newData := model.Faucet{
		Id:             primitive.NewObjectID(),
		Chain:          f.Chain,
		Public_address: f.Public_address,
		Last_txn_Time:  time.Now().UTC().String(),
		Txn_count:      f.Txn_count,
	}

	result, err := faucetCollection.InsertOne(ctx, newData)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(FaucetResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(FaucetResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func tokenTransfer(chain string, address string) {
	if chain == "ARV" {
		fmt.Println("arweave")
	} else if chain == "SOL" {
		fmt.Println("Solana")
	} else if chain == "DOJ" {
		fmt.Println("Dojima")
	} else {
		fmt.Println("Not available")
	}
}

// Faucet token transfer

func SendToken(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// var f model.Faucet
	defer cancel()

	fmt.Println(c.Params("chain"))
	fmt.Println(c.Params("address"))

	//use the validator library to validate required fields
	// if validationErr := validate.Struct(&f); validationErr != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(FaucetResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	// }

	cursor, err := faucetCollection.Find(ctx, bson.D{{"chain", c.Params("chain")}, {"public_address", c.Params("address")}})

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Catchphrases Not found",
			"error":   err,
		})
	}

	var blockData []model.Faucet

	for cursor.Next(ctx) {
		var b model.Faucet
		cursor.Decode(&b)
		blockData = append(blockData, b)
	}

	sort.SliceStable(blockData, func(i, j int) bool {
		return blockData[i].Last_txn_Time > blockData[j].Last_txn_Time
	})

	fmt.Println(blockData[0].Last_txn_Time)

	firstDate, _ := time.Parse(time.RFC3339, blockData[0].Last_txn_Time)
	secondDate := time.Now().UTC()
	difference := secondDate.Sub(firstDate)

	fmt.Println(difference)

	// if int64(difference) < 24 {
	// 	return c.JSON(fiber.Map{"status": "failed", "message": "Last transaction happened less than 24hrs ago.", "data": nil})
	// } else {
	// 	tokenTransfer(c.Params("chain"), c.Params("address"))
	// }
	tokenTransfer(c.Params("chain"), c.Params("address"))
	return c.JSON(fiber.Map{"status": "success", "message": "Faucet data Found from mongoDB", "data": blockData})

	return nil

}
