package runalyze

// Error provides access to the error message on returned errors
type Error struct {
	Message string `json:"message"`
}

func (e Error) Error() string { return string(e.Message) }

// AuthError indicates an issue with the authentication token
type AuthError string

func (e AuthError) Error() string { return string(e) }