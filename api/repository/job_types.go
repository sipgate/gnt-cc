package repository

import (
	"encoding/json"
	"fmt"
)

type rapiRemoteImportDisk struct {
	IpAddress  string
	Port       int
	Magic      string
	HmacDigest string
	Salt       string
}

func (r *rapiOpLogEntry) parse(b []byte) error {
	tmp := []interface{}{&r.Serial, &r.Timestamps, &r.PayloadType, &r.Payload}

	return unmarshalArrayIntoStruct(&tmp, b)
}

func (r *rapiRemoteImportDisk) parse(b []byte) error {
	tmp := []interface{}{&r.IpAddress, &r.Port, &r.Magic, &r.HmacDigest, &r.Salt}

	return unmarshalArrayIntoStruct(&tmp, b)
}

func unmarshalArrayIntoStruct(tmp *[]interface{}, b []byte) error {
	wantLen := len(*tmp)
	if err := json.Unmarshal(b, tmp); err != nil {
		return err
	}
	if g, e := len(*tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields: %d != %d", g, e)
	}
	return nil
}

type rapiRemoteImportPayload struct {
	Disks []json.RawMessage `json:"disks"`
}

type rapiOpLogEntry struct {
	Serial      int
	Timestamps  []int
	PayloadType string
	Payload     json.RawMessage
}

type rapiJobResponse struct {
	Status     string              `json:"status"`
	ID         int                 `json:"id"`
	StartTs    []int               `json:"start_ts"`
	EndTs      []int               `json:"end_ts"`
	ReceivedTs []int               `json:"received_ts"`
	Summary    []string            `json:"summary"`
	OpStatus   []string            `json:"opstatus"`
	Ops        []interface{}       `json:"ops"`
	OpLog      [][]json.RawMessage `json:"oplog"`
}
