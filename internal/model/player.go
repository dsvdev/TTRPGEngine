package model

type Player struct {
	Id    int64   `bson:"_id"`
	State StateID `bson:"state_id"`
}
