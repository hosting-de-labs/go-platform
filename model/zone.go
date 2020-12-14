package model

type ZoneObject struct {
	ZoneConfig ZoneConfigObject `json:"zoneConfig,omitempty"`
	Records    []RecordObject   `json:"records,omitempty"`
}

type FindZonesResult struct {
	GenericResultMetadata

	Data []ZoneObject `json:"data,omitempty"`
}

type ZonesResult struct {
	Metadata Metadata        `json:"metadata,omitempty"`
	Response FindZonesResult `json:"response,omitempty"`

	Status   string
	Warnings []string
}
