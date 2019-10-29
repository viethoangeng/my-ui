package app

// ProtocolProperties contains the device connection information in key/value pair
type ProtocolProperties map[string]string

// DeviceShort is the simple information of the device
type DeviceShort struct {
	NameID   string                        `json:"name"`      // Unique name for identifying a device
	UserName map[string]ProtocolProperties `json:"protocols"` // Description
}

// Reading struct will return for Cloud
type Reading struct {
	// Id       string `json:"id"`
	UserName string `json:"username"` // the name which the user will use
	Created  int64  `json:"created"`  // When the reading was created
	Device   string `json:"device"`   // the name was used in edgex
	Name     string `json:"name"`     // the name of Readings
	Value    string `json:"value"`    // Device sensor data value
}

//ValueDescriptor Struct will return Cloud
type ValueDescriptor struct {
	// Id            *string      `json:"id,omitempty"`
	Created     int64   `json:"created,omitempty"`
	Description *string `json:"description,omitempty"`
	// Modified      int64        `json:"modified,omitempty"`
	// Origin        int64        `json:"origin,omitempty"`
	Name          *string      `json:"name,omitempty"`
	Min           *interface{} `json:"min,omitempty"`
	Max           *interface{} `json:"max,omitempty"`
	DefaultValue  *interface{} `json:"defaultValue,omitempty"`
	Type          *string      `json:"type,omitempty"`
	UomLabel      *string      `json:"uomLabel,omitempty"`
	Formatting    *string      `json:"formatting,omitempty"`
	Labels        []string     `json:"labels,omitempty"`
	MediaType     *string      `json:"mediaType,omitempty"`
	FloatEncoding *string      `json:"floatEncoding,omitempty"`
}
