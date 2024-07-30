package requests

// ForgotPasswordRequest represents a request to reset a user's password.
type ForgotPasswordRequest struct {
	Email string `json:"email"`
}
