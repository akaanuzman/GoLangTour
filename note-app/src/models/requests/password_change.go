package requests

type PasswordChangeRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
