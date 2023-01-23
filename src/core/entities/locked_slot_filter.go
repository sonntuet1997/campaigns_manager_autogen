package entities

type LockedSlotFilter struct {
	*PagingFilter
	Code string
	Name string
}
