package patreon

var (
	// MemberDefaultIncludes specifies default includes for Member.
	MemberDefaultIncludes = []string{"address", "campaign", "currently_entitled_tiers", "user"}

	// MemberFields is all fields in the Member Attributes struct
	MemberFields = getObjectFields(Member{}.Attributes)
)

// Member is the record of a user's membership to a campaign.
// Remains consistent across months of pledging.
type Member struct {
	Type          string           `json:"type"`
	ID            string           `json:"id"`
	Attributes    MemberAttributes `json:"attributes"`
	Relationships struct {
		Address                *AddressRelationship     `json:"address,omitempty"`
		Campaign               *CampaignRelationship    `json:"campaign,omitempty"`
		CurrentlyEntitledTiers *TiersRelationship       `json:"currently_entitled_tiers,omitempty"`
		PledgeHistory          *PledgeEventRelationship `json:"pledge_history,omitempty"`
		User                   *UserRelationship        `json:"user,omitempty"`
	} `json:"relationships"`
}

// MemberAttributes is the attributes struct for Member
type MemberAttributes struct {
	CampaignLifetimeSupportCents int      `json:"campaign_lifetime_support_cents"`
	CurrentlyEntitledAmountCents int      `json:"currently_entitled_amount_cents"`
	Email                        string   `json:"email"`
	FullName                     string   `json:"full_name"`
	IsFollower                   bool     `json:"is_follower"`
	LastChargeDate               NullTime `json:"last_charge_date"`
	LastChargeStatus             string   `json:"last_charge_status"`
	LifetimeSupportCents         int      `json:"lifetime_support_cents"`
	NextChargeDate               NullTime `json:"next_charge_date"`
	Note                         string   `json:"note"`
	PatronStatus                 string   `json:"patron_status"`
	PledgeCadence                int      `json:"pledge_cadence"`
	PledgeRelationshipStart      NullTime `json:"pledge_relationship_start"`
	WillPayAmountCents           int      `json:"will_pay_amount_cents"`
}

type WebhookMemberAttributes struct {
	AccessExpiresAt              interface{} `json:"access_expires_at"`
	CampaignCurrency             string      `json:"campaign_currency"`
	CampaignLifetimeSupportCents int         `json:"campaign_lifetime_support_cents"`
	CampaignPledgeAmountCents    int         `json:"campaign_pledge_amount_cents"`
	FullName                     string      `json:"full_name"`
	IsFollower                   bool        `json:"is_follower"`
	IsFreeMember                 interface{} `json:"is_free_member"`
	IsFreeTrial                  interface{} `json:"is_free_trial"`
	LastChargeDate               NullTime   `json:"last_charge_date"`
	LastChargeStatus             string      `json:"last_charge_status"`
	LifetimeSupportCents         int         `json:"lifetime_support_cents"`
	PatronStatus                 string      `json:"patron_status"`
	PledgeAmountCents            int         `json:"pledge_amount_cents"`
	PledgeRelationshipStart      NullTime   `json:"pledge_relationship_start"`
}

// MemberResponse wraps Patreon's fetch benefit API response
type MemberResponse struct {
	Data     Member   `json:"data"`
	Included Includes `json:"included"`
	Links    struct {
		Self string `json:"self"`
	} `json:"links"`
}

// MembersResponse wraps Patreon's fetch benefit API response
type MembersResponse struct {
	Data     []Member `json:"data"`
	Included Includes `json:"included"`
	Links    struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Next  string `json:"next"`
		Prev  string `json:"prev"`
	} `json:"links"`
	Meta struct {
		Pagination struct {
			Cursors struct {
				First string `json:"first"`
				Last  string `json:"last"`
				Next  string `json:"next"`
				Prev  string `json:"prev"`
			} `json:"cursors"`
		} `json:"pagination"`
		Count int `json:"count"`
	} `json:"meta"`
}
