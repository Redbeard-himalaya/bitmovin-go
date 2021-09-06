package models

type EncodingScheduling struct {
	Priority *int64                  `json:"priority,omitempty"`
	PrewarmedEncoderPoolIds []string `json:"prewarmedEncoderPoolIds,omitempty"`
}
