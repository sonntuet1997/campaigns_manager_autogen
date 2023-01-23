package entities

type CampaignFilter struct {
	*PagingFilter
	Code string
	Name string
}
