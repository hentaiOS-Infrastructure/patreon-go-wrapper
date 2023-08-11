package patreon

const (
	// MemberCreate specifies a create pledge event
	MemberCreate = "members:create"

	// MemberUpdate specifies an update pledge event
	MemberUpdate = "members:update"

	// MemberDelete specifies a delete pledge event
	MemberDelete = "members:delete"

	// MemberCreatePledge specifies a create pledge event
	MemberCreatePledge = "members:pledge:create"

	// MemberUpdatePledge specifies an update pledge event
	MemberUpdatePledge = "members:pledge:update"

	// MemberDeletePledge specifies a delete pledge event
	MemberDeletePledge = "members:pledge:delete"
)


var (
	// WebhookDefaultIncludes specifies default includes for Webhook.
	WebhookDefaultIncludes = []string{"campaign", "client"}

	// WebhookFields is all fields in the Webhook Attributes struct
	WebhookFields = getObjectFields(Webhook{}.Attributes)
)

// Webhook is fired based on events happening on a particular campaign.
type Webhook struct {
	Type          string            `json:"type"`
	ID            string            `json:"id"`
	Attributes    WebhookAttributes `json:"attributes"`
	Relationships struct {
		Campaign    *CampaignRelationship    `json:"campaign,omitempty"`
		Memberships *MembershipsRelationship `json:"memberships,omitempty"`
	} `json:"relationships"`
}

// WebhookAttributes is the attributes struct for Webhook
type WebhookAttributes struct {
	LastAttemptedAt           NullTime    `json:"last_attempted_at"`
	NumConsecutiveTimesFailed int         `json:"num_consecutive_times_failed"`
	Paused                    bool        `json:"paused"`
	Secret                    string      `json:"secret"`
	Triggers                  interface{} `json:"triggers"`
	URI                       string      `json:"uri"`
}

// WebhookResponse wraps Patreon's fetch user API response
type WebhookResponse struct {
	Data     Webhook  `json:"data"`
	Included Includes `json:"included"`
	Links    struct {
		Self string `json:"self"`
	} `json:"links"`
}
