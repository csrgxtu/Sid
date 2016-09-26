package services

import (
  "github.com/astaxie/beego"
  "gopkg.in/mgo.v2/bson"
  "Sid/models"
  "errors"
)

var UserCollection = beego.AppConfig.String("UserCollection")
var VirusCollection = beego.AppConfig.String("VirusCollection")
var InfectedCollection = beego.AppConfig.String("InfectedCollection")

func GetVirus(vid string) (err error, rtv models.Virus) {
  if CheckAndReconnect() != nil {
    return
  }

  var criteria = bson.M{"vid": vid}
  err = Session.DB(DB).C(VirusCollection).Find(criteria).One(&rtv)
  if err != nil {
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  return
}

func GetInfects(vid string) (err error, rtv []models.Infected) {
  if CheckAndReconnect() != nil {
    return
  }

  var criteria = bson.M{"vid": vid}
  err = Session.DB(DB).C(InfectedCollection).Find(criteria).All(&rtv)
  if err != nil {
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  return
}

// var ScanImageCollection = beego.AppConfig.String("ScanImageCollection")
// var SpineImageCollection = beego.AppConfig.String("SpineImageCollection")

// func SearchScanImage(user_id string, skip, limit int) (err error, rtv []models.ScanImage) {
//   if CheckAndReconnect() != nil {
//     return
//   }
//
//   var criteria = bson.M{"user_id": user_id}
//   err = Session.DB(DB).C(ScanImageCollection).Find(criteria).Skip(skip).Limit(limit).Sort("-createtime").All(&rtv)
//   if err != nil {
//     beego.Info(err)
//     err = errors.New("Server Internal Error")
//     return
//   }
//
//   return
// }
//
// func ReadScanImage(skip, limit int) (err error, rtv []models.ScanImage) {
//   if CheckAndReconnect() != nil {
//     return
//   }
//
//   var criteria = bson.M{"status": "visable"}
//   err = Session.DB(DB).C(ScanImageCollection).Find(criteria).Sort("-createtime").Skip(skip).Limit(limit).All(&rtv)
//   if err != nil {
//     err = errors.New("Server Internal Error")
//     return
//   }
//   return
// }
//
