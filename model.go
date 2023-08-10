package db

type UserProfile struct {
	Username      string `json:"username"` // unique per user
	FullName      string `json:"full_name"`
	Bio           string `json:"bio"`
	ProfilePicURL string `json:"profile_pic_url"`
}
