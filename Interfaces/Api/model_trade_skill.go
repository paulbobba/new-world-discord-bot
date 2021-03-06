/*
 * New World Trade Skill Tracker
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TradeSkill struct {
	discordId primitive.ObjectID `bson:"_id,omitempty"`
	Gathering *Gathering         `json:"gathering,omitempty" bson:"gathering,omitempty"`
	Refining  *Refining          `json:"refining,omitempty" bson:"refining,omitempty"`
	Crafting  *Crafting          `json:"crafting,omitempty" bson:"crafting,omitempty"`
}
