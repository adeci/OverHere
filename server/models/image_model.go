package models

type Image struct {
	ImageID  string  `json:"imageid,omitempty"`
	UserID   string  `json:"userid" validate:"required"`
	OHPostID string  `json:"ohpostid,omitempty"`
	Encoding string  `json:"encoding" validate:"required"`
	XCoord   float64 `json:"xcoord" validate:"required"`
	YCoord   float64 `json:"ycoord" validate:"required"`
}
