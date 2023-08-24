package model

type UserProfile struct {
	ID            int
	Username      string `json:"username"` // unique per user
	FullName      string `json:"full_name"`
	Bio           string `json:"bio"`
	ProfilePicURL string `json:"profile_pic_url"`
}

// check it's supposed to be in db package
