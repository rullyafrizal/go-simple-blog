package requests

type StorePostRequest struct {
	Title       string `form:"title"`
	Content     string `form:"content"`
	CategoryId  uint64 `form:"category_id"`
}
