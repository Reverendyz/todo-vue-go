package types

type Document struct {
	Task string `bson:"task"`
	Done bool   `bson:"done"`
}
