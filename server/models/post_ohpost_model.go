package models

type PostOHPost struct {
	UserID   string   `json:"userid"`
	Tag      string   `json:"tag,omitempty"`
	Caption  string   `json:"caption,omitempty"`
	ImageIds []string `json:"imageids"`
}
