package delivery

import (
	"fmt"
	"net/http"
	"religion/config"
	"religion/internal/domain/user"
	"religion/internal/domain/user/dto"
	"religion/pkg/httperrors"
	"religion/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type userHandlers struct {
	cfg        *config.Config
	userGroup  *echo.Group
	groupGroup *echo.Group
	userUC     user.UseCase
	logger     logger.Logger
	validate   *validator.Validate
}

func NewUserHandlers(
	cfg *config.Config,
	group *echo.Group,
	ggroup *echo.Group,
	userUC user.UseCase,
	logger logger.Logger,
	validate *validator.Validate,
) *userHandlers {
	return &userHandlers{cfg: cfg, userGroup: group, groupGroup: ggroup, userUC: userUC, logger: logger, validate: validate}
}

func (h *userHandlers) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var input dto.CreateUserRequest

		if err := c.Bind(&input); err != nil {
			h.logger.Errorf("c.Bind: %v", err)
			return httperrors.ErrorCtxResponse(c, err)
		}

		if err := h.validate.StructCtx(ctx, &input); err != nil {
			h.logger.Errorf("validate.StructCtx: %v", err)
			return httperrors.ErrorCtxResponse(c, err)
		}

		user, err := h.userUC.CreateUser(ctx, &input)
		if err != nil {
			h.logger.Errorf("userUC.CreateUser: %v", err)
			return httperrors.ErrorCtxResponse(c, err)
		}

		return c.JSON(http.StatusCreated, &user)
	}
}

func (h *userHandlers) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var input dto.UpdateUserRequest

		if err := c.Bind(&input); err != nil {
			h.logger.Errorf("c.Bind: %v", err)
			return httperrors.ErrorCtxResponse(c, err)
		}

		if err := h.validate.StructCtx(ctx, &input); err != nil {
			h.logger.Errorf("validate.StructCtx: %v", err)
			return httperrors.ErrorCtxResponse(c, err)
		}

		user, err := h.userUC.UpdateUser(ctx, &input)
		if err != nil {
			h.logger.Errorf("userUC.UpdateUser: %v", err)
			return httperrors.ErrorCtxResponse(c, err)
		}

		return c.JSON(http.StatusAccepted, &user)
	}
}

func (h *userHandlers) GetSingleUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		//uid := c.Param("uid")
		uid := c.Get("token").(string)
		h.logger.Infof("gets uid %s", uid)
		if uid == "" {
			h.logger.Errorf("ctx.Value: %v", "uid empty")
			return httperrors.ErrorCtxResponse(c, fmt.Errorf("empty uid"))
		}

		user, err := h.userUC.GetSingleUser(ctx, uid)
		if err != nil {
			h.logger.Errorf("userUC.GetSingleUser: %v", err)
			return httperrors.ErrorCtxResponse(c, err)
		}

		return c.JSON(http.StatusOK, &user)
	}
}

func (h *userHandlers) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		users, err := h.userUC.GetAllUsers(ctx)
		if err != nil {
			h.logger.Errorf("userUC.GetAllUsers: %v", err)
			return httperrors.ErrorCtxResponse(c, err)
		}

		return c.JSON(http.StatusAccepted, &users)
	}
}

func (h *userHandlers) CreateGroup() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var input dto.GroupRequest

		if err := c.Bind(&input); err != nil {
			h.logger.Errorf("c.Bind: %v", err)
			return httperrors.ErrorCtxResponse(c, err)
		}

		if err := h.validate.StructCtx(ctx, &input); err != nil {
			h.logger.Errorf("validate.StructCtx: %v", err)
			return httperrors.ErrorCtxResponse(c, err)
		}

		group, err := h.userUC.CreateGroup(ctx, &input)
		if err != nil {
			h.logger.Errorf("userUC.CreateGroup: %v", err)
			return httperrors.ErrorCtxResponse(c, err)
		}

		return c.JSON(http.StatusCreated, &group)
	}
}
