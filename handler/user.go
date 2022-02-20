package handler

import (
	"Onboarding/business"
	"Onboarding/business/directus"
	"Onboarding/utils"
	"github.com/labstack/echo/v4"
	"strings"
)

func UseUserHandler(e *echo.Echo) {
	e.POST("/user/check_external", checkUserExternal)
	e.POST("/user/login_external", requestTokenFromExternal)
	e.POST("/user/create", checkUserJWT(createUser))
}

type UserAuth struct {
	Username string `json:"username" form:"username" validate:"required,max=12"`
	Password string `json:"password" form:"password" validate:"required,min=4,max=32"`
}

// checkUserExternal godoc
// @Summary 检查我的珠科用户名密码是否有效
// @Description 检查我的珠科账号
// @Tags 用户
// @Accept json
// @Produce json
// @Success 204 "响应体为空"
// @Failure 401 "用户名或密码错误/服务器连接错误"
// @Router /user/check_external [post]
func checkUserExternal(c echo.Context) error {
	tokens := new(UserAuth)
	if err := c.Bind(tokens); err != nil {
		return err
	}
	if err := c.Validate(tokens); err != nil {
		return c.JSON(400, utils.MessageResponse{Message: "Bad Request"})
	}

	if business.CheckExternalUser(tokens.Username, tokens.Password) {
		return c.String(204, "")
	} else {
		return c.JSON(401, map[string]interface{}{
			"message": "Error Certification",
		})
	}
}

// requestTokenFromExternal godoc
// @Summary 使用我的珠科账号密码换取凭据
// @Description 如果成功，则会返回一个JWT。若返回的是空字符串，说明服务端配置错误。
// @Tags 用户
// @Param username body string true "我的珠科用户名"
// @Param password body string true "我的珠科密码"
// @Accept json
// @Produce json
// @Success 201 {object} utils.TokenResponse
// @Failure 401 {object} utils.MessageResponse
// @Router /user/login_external [post]
func requestTokenFromExternal(c echo.Context) error {
	tokens := new(UserAuth)
	if err := c.Bind(tokens); err != nil {
		return err
	}
	if err := c.Validate(tokens); err != nil {
		return c.JSON(400, utils.MessageResponse{Message: "Bad Request"})
	}
	if business.CheckExternalUser(tokens.Username, tokens.Password) {
		token := utils.NewJWT("On-boarding Token", &utils.AnyStruct{
			"username": tokens.Username,
		})
		return c.JSON(201, utils.TokenResponse{
			MessageResponse: utils.MessageResponse{
				Message: "Created",
			},
			Token: token,
		})
	} else {
		return c.JSON(401, utils.MessageResponse{
			Message: "Error Certification",
		})
	}
}

func checkUserJWT(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if !strings.Contains(header, "Bearer ") {
			return c.JSON(401, utils.MessageResponse{Message: "No Authorization"})
		}
		header = strings.Replace(header, "Bearer ", "", 1)
		claims, err := utils.CheckJWT(header)
		if err != nil {
			return c.JSON(401, utils.MessageResponse{Message: "Bad JSON Web Token"})
		}
		c.Set("JWT", claims)
		return h(c)
	}
}

// createUser godoc
// @Summary 创建OA(Directus)用户
// @Description 在换取完凭据后，通过Bearer Token传递凭据，可以创建一个用户。
// @Tags 用户
// @Param Authorization header string true "Bearer <JWT>"
// @Param email body string true "登录邮箱"
// @Param password body string true "登录密码"
// @Accept json
// @Produce json
// @Success 201 {object} utils.MessageResponse
// @Failure 401 {object} utils.MessageResponse
// @Failure 409 {object} utils.MessageResponse
// @Router /user/create [post]
func createUser(c echo.Context) error {
	user := new(directus.BaseUser)
	if err := c.Bind(user); err != nil {
		return err
	}
	if err := c.Validate(user); err != nil {
		return c.JSON(400, utils.MessageResponse{Message: "Bad Request"})
	}

	userFilled := directus.User{
		BaseUser: *user,
		Role:     "31f7ba59-e24b-4992-96a9-f1c24bd22d99",
	}

	if userFilled.Create() {
		return c.JSON(201, utils.MessageResponse{Message: "Created"})
	} else {
		return c.JSON(409, utils.MessageResponse{Message: "Conflicted"})
	}
}
