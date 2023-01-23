package models

type LockedSlot struct {
	ID   uint64
	Code string
	Name string
}

func (LockedSlot) TableName() string {
	return "locked-slots"
}
