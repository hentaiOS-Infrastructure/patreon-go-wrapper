package patreon

// RewardFields is all fields in the Reward Attributes struct
var RewardFields = getObjectFields(Reward{}.Attributes)

// Reward is a membership level on a patron, which can have benefits attached to it.
type Reward struct {
	Type          string           `json:"type"`
	ID            string           `json:"id"`
	Attributes    RewardAttributes `json:"attributes"`
	Relationships struct {
		Campaign *CampaignRelationship `json:"campaign,omitempty"`
	} `json:"relationships"`
}

type RewardAttributes struct {
	Amount              int         `json:"amount"`
	AmountCents         int         `json:"amount_cents"`
	CreatedAt           NullTime    `json:"created_at"`
	Currency            string      `json:"currency"`
	DeclinedPatronCount int         `json:"declined_patron_count"`
	Description         string      `json:"description"`
	DiscordRoleIds      interface{} `json:"discord_role_ids"`
	EditedAt            NullTime    `json:"edited_at"`
	ImageURL            interface{} `json:"image_url"`
	IsFreeTier          bool        `json:"is_free_tier"`
	PatronAmountCents   int         `json:"patron_amount_cents"`
	PatronCount         int         `json:"patron_count"`
	PatronCurrency      string      `json:"patron_currency"`
	PostCount           int         `json:"post_count"`
	Published           bool        `json:"published"`
	PublishedAt         NullTime    `json:"published_at"`
	Remaining           interface{} `json:"remaining"`
	RequiresShipping    bool        `json:"requires_shipping"`
	Title               string      `json:"title"`
	UnpublishedAt       interface{} `json:"unpublished_at"`
	URL                 string      `json:"url"`
	UserLimit           interface{} `json:"user_limit"`
}
