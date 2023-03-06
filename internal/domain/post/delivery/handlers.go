package delivery

import (
	"net/http"
	"religion/config"
	"religion/internal/domain/post"
	"religion/pkg/httperrors"
	"religion/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type postHandlers struct {
	cfg      *config.Config
	group    *echo.Group
	postUC   post.UseCase
	logger   logger.Logger
	validate *validator.Validate
}

func NewPostHandlers(
	cfg *config.Config,
	group *echo.Group,
	postUC post.UseCase,
	logger logger.Logger,
	validate *validator.Validate,
) *postHandlers {
	return &postHandlers{cfg: cfg, group: group, postUC: postUC, logger: logger, validate: validate}
}

func (h *postHandlers) GetPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		uid := c.Get("token").(string)

		posts, err := h.postUC.GetPosts(ctx, uid)
		if err != nil {
			h.logger.Errorf("postsUC.GetPosts: %v", err)
			return httperrors.ErrorCtxResponse(c, err)
		}

		return c.JSON(http.StatusAccepted, &posts)
	}
}
