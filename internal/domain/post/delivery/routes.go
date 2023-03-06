package delivery

func (h *postHandlers) MapPostsRoutes() {
	h.group.GET("/all", h.GetPosts())
}
