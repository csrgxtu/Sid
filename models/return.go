package models

type (
  Recs interface {}

	// Result struct {
	// 	Msg   string `json:"msg" bson:"msg"`
	// 	Data  []Recs `json:"data" bson:"data"`
	// }

  Result struct {
    Msg string `json:"msg"`
    Data VirusInfectedUsers `json:"data"`
  }

  VirusInfectedUsers struct {
    Virus Virus `json:"virus"`
    Infected []Infected `json:"infected"`
    Users []User `json:"user"`
  }

  D3 struct {
    Name string `json:"name" bson:"name"`
    Children []D3 `json:"children" bson:"children"`
  }
)
