package base

import (
	"encoding/json"
	"fmt"
)

type BluetoothRing struct {
	TagMac      string `json:"tag_mac"`
	Rssi        int8   `json:"rssi"`
	RingMac     string `json:"ring_mac"`
	DegreeX     int8   `json:"degree_x"`
	DegreeY     int8   `json:"degree_y"`
	DegreeZ     int8   `json:"degree_z"`
	Bett        uint8  `json:"bett"`
	Warn        uint8  `json:"warn"`
	X           int32  `json:"x"`
	Y           int32  `json:"y"`
	Orientation string `json:"orientation"`
	FingerID    string `json:"finger_id"`
}

func (btr *BluetoothRing) UnmarshalJSON(str string) error {
	tmp := []interface{}{&btr.TagMac, &btr.Rssi, &btr.RingMac, &btr.DegreeX, &btr.DegreeY, &btr.DegreeZ, &btr.Bett, &btr.Warn, &btr.X, &btr.Y, &btr.Orientation, &btr.FingerID}
	wantLen := len(tmp)
	if err := json.Unmarshal([]byte(str), &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Notification: %d != %d", g, e)
	}
	return nil
}
