package model

type UserProfile struct {
	ID            int    `json:"userID"`          // ID is the auto-generated unique identifier for the user's profile
	Username      string `json:"username"`        // Username is the unique username per user
	FullName      string `json:"full_name"`       // FullName is the user's full name
	Bio           string `json:"bio"`             // Bio is the user's profile biography
	ProfilePicURL string `json:"profile_pic_url"` // ProfilePicURL is the URL to the user's profile picture
}

// todo: check if i need the id and the tag
