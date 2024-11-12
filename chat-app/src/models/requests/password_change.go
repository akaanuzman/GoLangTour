package requests

// PasswordChangeRequest represents a request to change a user's password.
type PasswordChangeRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
