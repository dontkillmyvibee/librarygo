package types

import "encoding/json"

type Nullable[T any] struct {
	Value *T
	Set   bool
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	var value *T

	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	n.Value = value
	n.Set = true

	return nil
}
