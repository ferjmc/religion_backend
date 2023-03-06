package dto

type GroupRequest struct {
	UserUid string `validate:"required"`
	Name    string
}
