package structs

type EmailAndPasswordReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResp struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
}

type GetEntryResp struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateEntryReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type DeleteEntryReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Title    string `json:"title"`
}
