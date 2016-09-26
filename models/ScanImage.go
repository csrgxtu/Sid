package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	ScanImage struct {
			Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
			UserId string `json:"user_id" bson:"user_id"`
			Image string `json:"image" bson:"image"`
			Location []string `json:"location" bson:"location"`
			SpineCount int `json:"spine_count" bson:"spine_count"`
			Favour int `json:"favour" bson:"favour"`
			Count int `json:"count" bson:"count"`
			State string `json:"state" bson:"state"` // 在录书和申书操作中的状态, [CheckPass, CheckIgnore, CheckRM, ValidatePass, ValidateIgnore, ValidateRM]
			Status string `json:"status" bson:"status"`
			Version time.Time `json:"version" bson:"version"`
			CreateUserId string `json:"createuserid" bson:"createuserid"`
			UpdateUserId string `json:"updateuserid" bson:"updateuserid"`
			CreateTime string `json:"createtime" bson:"createtime"`
			UpdateTime string `json:"updatetime" bson:"updatetime"`
	}

	SpineImage struct {
		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Url string `json:"url" bson:"url"`
		Sources string `json:"sources" bson:"sources"`
		Isbn string `json:"isbn" bson:"isbn"`
		Type string `json:"type" bson:"type"`
		FloorId string `json:"floor_id" bson:"floor_id"`
		ScanId string `json:"scan_id" bson:"scan_id"`
		Bid string `json:"bid" bson:"bid"`
		Meta []string `json:"meta" bson:"meta"`
		UserVipLevel int `json:"user_vip_level" bson:"user_vip_level"`
		IsPush bool `json:"isPush" bson:"isPush"`
		State string `json:"state" bson:"state"` // 在录书和申书操作中的状态, [CheckPass, CheckIgnore, CheckRMValidatePass, ValidateIgnore, ValidateRM]
		Status string `json:"status" bson:"status"`
		Version time.Time `json:"version" bson:"version"`
		CreateUserId string `json:"createuserid" bson:"createuserid"`
		UpdateUserId string `json:"updateuserid" bson:"updateuserid"`
		CreateTime string `json:"createtime" bson:"createtime"`
		UpdateTime string `json:"updatetime" bson:"updatetime"`
	}
)
