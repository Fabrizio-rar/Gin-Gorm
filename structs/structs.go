package structs

type EmailAndPasswordReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateEntryReq struct {
	Email   string `json:"email"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetUserResp struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
}
