package swagger

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	db "NewWorldBot/Interfaces/TradeSkills/DataAccess"
)

func TradeSkillUpsert(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var tradeSkill *TradeSkill
	err := decoder.Decode(&tradeSkill)

	fmt.Printf("%#v\n", *tradeSkill)

	if err != nil {
		fmt.Println(err)

		return
	}

	params := mux.Vars(r)
	discordId, _ := params["discordId"]
	(*tradeSkill).discordId, _ = primitive.ObjectIDFromHex(discordId)
	tradeSkillDoc := &TradeSkillFlattened{
		DiscordId:              (*tradeSkill).discordId,
		GatheringLogging:       (*tradeSkill).Gathering.Logging,
		GatheringMining:        (*tradeSkill).Gathering.Mining,
		GatheringFishing:       (*tradeSkill).Gathering.Fishing,
		GatheringHarvesting:    (*tradeSkill).Gathering.Harvesting,
		GatheringSkinning:      (*tradeSkill).Gathering.Skinning,
		RefiningSmelting:       (*tradeSkill).Refining.Smelting,
		RefiningWoodworking:    (*tradeSkill).Refining.Woodworking,
		RefiningLeatherworking: (*tradeSkill).Refining.Leatherworking,
		RefiningWeaving:        (*tradeSkill).Refining.Weaving,
		RefiningStonecutting:   (*tradeSkill).Refining.Stonecutting,
		CraftingWeaponsmithing: (*tradeSkill).Crafting.Weaponsmithing,
		CraftingArmoring:       (*tradeSkill).Crafting.Armoring,
		CraftingEngineering:    (*tradeSkill).Crafting.Engineering,
		CraftingJewelcrafting:  (*tradeSkill).Crafting.Jewelcrafting,
		CraftingArcana:         (*tradeSkill).Crafting.Arcana,
		CraftingCooking:        (*tradeSkill).Crafting.Cooking,
		CraftingFurnishing:     (*tradeSkill).Crafting.Furnishing,
	}
	mongoDBConnection := os.Getenv(db.MongoDBConnectionStringEnvVarName)
	mongoClient := db.NewClient(mongoDBConnection)
	mongoDBDatabase := os.Getenv(db.MongoDBDatabaseStringEnvVarName)
	mongoDBCollection := os.Getenv(db.MongoDBCollectionStringEnvVarName)
	collection := mongoClient.Database(mongoDBDatabase).Collection(mongoDBCollection)

	filter := bson.M{"_id": bson.M{"$eq": discordId}}
	update := bson.M{"$set": tradeSkillDoc}
	fmt.Println("update", update)
	opts := options.Update().SetUpsert(true)

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	result, err := collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		fmt.Println("UpdateOne() result ERROR:", err)
		os.Exit(1)
	} else {
		fmt.Println("UpdateOne() result:", result)
		fmt.Println("UpdateOne() result TYPE:", reflect.TypeOf(result))
		fmt.Println("UpdateOne() result MatchedCount:", result.MatchedCount)
		fmt.Println("UpdateOne() result ModifiedCount:", result.ModifiedCount)
		fmt.Println("UpdateOne() result UpsertedCount:", result.UpsertedCount)
		fmt.Println("UpdateOne() result UpsertedID:", result.UpsertedID)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

type TradeSkillFlattened struct {
	DiscordId              primitive.ObjectID `bson:"_id,omitempty"`
	GatheringLogging       int32              `bson:"gathering.logging,omitempty"`
	GatheringMining        int32              `bson:"gathering.mining,omitempty"`
	GatheringFishing       int32              `bson:"gathering.fishing,omitempty"`
	GatheringHarvesting    int32              `bson:"gathering.harvesting,omitempty"`
	GatheringSkinning      int32              `bson:"gathering.skinning,omitempty"`
	RefiningSmelting       int32              `bson:"refining.smelting,omitempty"`
	RefiningWoodworking    int32              `bson:"refining.woodworking,omitempty"`
	RefiningLeatherworking int32              `bson:"refining.leatherworking,omitempty"`
	RefiningWeaving        int32              `bson:"refining.weaving,omitempty"`
	RefiningStonecutting   int32              `bson:"refining.stonecutting,omitempty"`
	CraftingWeaponsmithing int32              `bson:"crafting.weaponsmithing,omitempty"`
	CraftingArmoring       int32              `bson:"crafting.armoring,omitempty"`
	CraftingEngineering    int32              `bson:"crafting.engineering,omitempty"`
	CraftingJewelcrafting  int32              `bson:"crafting.jewelcrafting,omitempty"`
	CraftingArcana         int32              `bson:"crafting.arcana,omitempty"`
	CraftingCooking        int32              `bson:"crafting.cooking,omitempty"`
	CraftingFurnishing     int32              `bson:"crafting.furnishing,omitempty"`
}
