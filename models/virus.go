package models

import (
	"gopkg.in/mgo.v2/bson"
	// "time"
)

type (
  Virus struct {
    Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Vid string `json:"vid" bson:"vid"`
    Url string `json:"url" bson:"url"`
    UserId string `json:"userid" bson:"userid"`
    Content string `json:"content" bson:"content"`
    Type string `json:"type" bson:"type"`
    // CreateTime time.Time `json:"createtime" bson:"createtime"`
    CreateTime int64 `json:"createtime" bson:"createtime"`
  }

  User struct {
    Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
    OpenId string `json:"openid" bson:"openid"`
    NickName string `json:"nickname" bson:"nickname"`
    Sex bool `json:"sex" bson:"sex"`
    Language string `json:"language" bson:"language"`
    City string `json:"city" bson:"city"`
    Province string `json:"province" bson:"province"`
    Country string `json:"country" bson:"country"`
    HeadImgUrl string `json:"headimgurl" bson:"headimgurl"`
    UnionId string `json:"unionid" bson:"unionid"`
    CreateTime int64 `json:"createtime" bson:"createtime"`
  }

  Infected struct {
    Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
    CarryId string `json:"carryid" bson:"carryid"`
    Vid string `json:"vid" bson:"vid"`
    InfectedId string `json:"infectid" bson:"infectid"`
    CreateTime int64 `json:"createtime" bson:"createtime"`
  }
)
