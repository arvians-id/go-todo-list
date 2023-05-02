package request

type ActivityCreateRequest struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type ActivityUpdateRequest struct {
	Title string `json:"title"`
}
