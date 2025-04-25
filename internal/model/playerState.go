package model

type StateID string

const (
	DefaultStateID StateID = "default"
	MedicStateID   StateID = "medic"
	CaptainStateID StateID = "captain"
)

type PlayerState struct {
	Id          StateID            `bson:"_id"`
	Image       string             `bson:"image"`
	Text        string             `bson:"text"`
	Transitions map[string]StateID `bson:"transitions"`
}
