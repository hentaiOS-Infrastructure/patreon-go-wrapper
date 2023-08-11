package patreon

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

const (
	// HeaderEventType specifies an event type HTTP header name
	HeaderEventType = "X-Patreon-Event"

	// HeaderEventSignature specifies message signature HTTP header name to verify message body
	HeaderSignature = "X-Patreon-Signature"
)

var (
	// WebhookTriggersDefaultIncludes specifies default includes for Webhook.
	WebhookTriggersDefaultIncludes = []string{"campaign", "user", "reward"}

	// WebhookTriggersFields is all fields in the Webhook Attributes struct
	// WebhookTriggers is pretty much Member Attributes.
	WebhookTriggersFields = getObjectFields(WebhookMember{}.Attributes)
)

// VerifySignature verifies the sender of the message
func VerifySignature(message []byte, secret string, signature string) (bool, error) {
	hash := hmac.New(md5.New, []byte(secret))
	if _, err := hash.Write(message); err != nil {
		return false, err
	}

	sum := hash.Sum(nil)
	expectedSignature := hex.EncodeToString(sum)

	return expectedSignature == signature, nil
}

type WebhookTriggersResponse struct {
	Data     WebhookMember `json:"data"`
	Included Includes      `json:"included"`
}
