package model

type ZoneObject struct {
	ZoneConfig ZoneConfigObject `json:"zoneConfig,omitempty"`
	Records    []RecordObject   `json:"records,omitempty"`
}
