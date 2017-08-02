package superego

// Profile represents a person's profile.
type Profile struct {
	// The ID of the profile
	ID string `json:"id", datastore:"-"`
	// The name of the person, which is suitable for display.
	DisplayName string `json:"displayName"`
	// A representation of the individual components of a person's name.
	Name Name `json:"name"`
	// The email address.
	Email string `json:"email"`
	// The URL of the person's profile photo.
	ImageURL string `json:"imageUrl"`
	// A short biography for this person.
	AboutMe string `json:"aboutMe"`
}

// Name represents the individual components of a person's name.
type Name struct {
	// The full name of this person.
	Formatted string `json:"formatted"`
	// The family name (last name) of this person.
	FamilyName string `json:"familyName"`
	// The given name (first name) of this person.
	GivenName string `json:"givenName"`
}
