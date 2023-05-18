package structs

type UpdateEntryReq struct {
	Email   string `json:"email"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
