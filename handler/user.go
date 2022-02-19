package handler

import (
    "Onboarding/business"
    "Onboarding/utils"
    "github.com/labstack/echo/v4"
)

func UseUserHandler(e *echo.Echo) {
    e.POST("/user/check_external", checkUserExternal)
    e.POST("/user/login_external", requestTokenFromExternal)
}

type UserAuth struct {
    Username string `json:"username" form:"username"`
    Password string `json:"password" form:"password"`
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
