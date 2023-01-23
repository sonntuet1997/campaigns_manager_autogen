package models

type Campaign struct {
	ID   uint64
	Code string
}

func (Campaign) TableName() string {
	return "campaigns"
}
