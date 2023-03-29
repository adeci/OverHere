package models

type OHPost struct {
	OHPostID string `json:"ohpostid,omitempty"`
	UserID   string `json:"userid"`
	Tag      string `json:"tag,omitempty"`
	Caption  string `json:"caption,omitempty"`
}
