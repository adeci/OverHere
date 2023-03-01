package models

type Image struct {
	ImageID  string `json:"imageid,omitempty"`
	UserID   string `json:"userid" validate:"required"`
	OHPostID string `json:"ohpostid,omitempty"`
	Encoding string `json:"encoding" validate:"required"`
}
