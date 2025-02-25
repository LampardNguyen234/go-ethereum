// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package catalyst

import (
	"encoding/json"

	"github.com/LampardNguyen234/go-ethereum/common/hexutil"
)

var _ = (*payloadResponseMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (p PayloadResponse) MarshalJSON() ([]byte, error) {
	type PayloadResponse struct {
		PayloadID hexutil.Uint64 `json:"payloadId"`
	}
	var enc PayloadResponse
	enc.PayloadID = hexutil.Uint64(p.PayloadID)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (p *PayloadResponse) UnmarshalJSON(input []byte) error {
	type PayloadResponse struct {
		PayloadID *hexutil.Uint64 `json:"payloadId"`
	}
	var dec PayloadResponse
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.PayloadID != nil {
		p.PayloadID = uint64(*dec.PayloadID)
	}
	return nil
}
