package delivery

func (h *userHandlers) MapUserRoutes() {
	h.userGroup.POST("/register", h.CreateUser())
	h.userGroup.POST("/update", h.UpdateUser())
	h.userGroup.GET("", h.GetSingleUser())
	h.userGroup.GET("/all", h.GetAllUsers())
}

func (h *userHandlers) MapGroupRoutes() {
	h.groupGroup.POST("/", h.CreateGroup())
}
