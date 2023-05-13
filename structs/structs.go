package structs

type EmailReq struct {
	Email string `json:"email"`
}

type TitleReq struct {
	Title string `json:"title"`
}

type UpdateEntryReq struct {
	Email   string `json:"email"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
