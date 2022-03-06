package requests

type StorePostRequest struct {
	Title       string `form:"title"`
	Content     string `form:"content"`
	Image       string `form:"image"`
	CategoryId  uint64 `form:"category_id"`
	IsPublished string `form:"is_published"`
}
