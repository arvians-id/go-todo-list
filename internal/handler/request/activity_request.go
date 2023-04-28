package request

type ActivityCreateRequest struct {
	Title string `json:"title" validate:"required,max=100"`
	Email string `json:"email" validate:"required,email,max=100"`
}

type ActivityUpdateRequest struct {
	Title string `json:"title" validate:"required,max=100"`
}
