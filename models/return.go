package models

type (
  Recs interface {}

	Result struct {
		Msg   string `json:"msg" bson:"msg"`
		Data  []Recs `json:"data" bson:"data"`
	}
)
