package delivery

func (h *userHandlers) MapUserRoutes() {
	h.group.POST("/register", h.CreateUser())
	h.group.POST("/update", h.UpdateUser())
	h.group.GET("", h.GetSingleUser())
	h.group.GET("/all", h.GetAllUsers())
}
