package patreon

import (
	"encoding/json"
	"fmt"
)

// Includes wraps 'includes' JSON field to handle objects of different type within an array.
type Includes struct {
	Items []interface{}
}

// UnmarshalJSON deserializes 'includes' field into the appropriate structs depending on the 'type' field.
// See http://gregtrowbridge.com/golang-json-serialization-with-interfaces/ for implementation details.
func (i *Includes) UnmarshalJSON(b []byte) error {
	var items []*json.RawMessage
	if err := json.Unmarshal(b, &items); err != nil {
		return err
	}

	count := len(items)
	i.Items = make([]interface{}, count)

	s := struct {
		Type string `json:"type"`
	}{}

	for idx, raw := range items {
		if err := json.Unmarshal(*raw, &s); err != nil {
			return err
		}

		var obj interface{}

		// Depending on the type, we can run json.Unmarshal again on the same byte slice
		// But this time, we'll pass in the appropriate struct instead of a map
		switch s.Type {
		case "user":
			obj = &User{}
		case "tier":
			obj = &Tier{}
		case "goal":
			obj = &Goal{}
		case "campaign":
			obj = &Campaign{}
		case "benefit":
			obj = &Benefit{}
		case "membership":
			obj = &Member{}
		case "member":
			obj = &Member{}
		case "address":
			obj = &Address{}
		case "patron":
			obj = &User{}
		case "webhook":
			obj = &Webhook{}
		case "deliverable":
			obj = &Deliverable{}
		case "reward":
			obj = &Reward{}
		default:
			return fmt.Errorf("unsupported type '%s'", s.Type)
		}

		if err := json.Unmarshal(*raw, obj); err != nil {
			return err
		}

		i.Items[idx] = obj
	}

	return nil
}
