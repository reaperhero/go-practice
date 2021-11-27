package http

import (
	"github.com/casbin/casbin/v2"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/reaperhero/casbin-echo-role/model"
	"net/http"
	"time"
)

type httpHander struct {
	users    model.Users
	enforcer *casbin.Enforcer
}

func NewHTTPHandler(e *echo.Echo, enforcer *casbin.Enforcer) {
	hander := &httpHander{
		users:    model.CreateUsers(),
		enforcer: enforcer,
	}
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, time.Now().Format(time.RFC3339Nano))
	})

	e.GET("/ping", hander.ping)
	e.POST("/login", hander.loginHandler)
	e.POST("/logout", hander.logoutHandler)

}

func (h *httpHander) ping(c echo.Context) error {
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "write" // the operation that the user performs on the resource.
	ok, _ := h.enforcer.Enforce(sub, obj, act)

	return c.JSON(200, ok)
}

func (h *httpHander) loginHandler(c echo.Context) error {
	name := c.Request().PostFormValue("name")
	user, err := h.users.FindByName(name)
	if err != nil {
		c.JSON(200, "用户不存在")
	}
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["userID"] = user.ID
	sess.Values["role"] = user.Role
	sess.Save(c.Request(), c.Response())
	return c.NoContent(http.StatusOK)
}

func (h *httpHander) logoutHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)
	if sess.IsNew {
		return c.JSON(200, "err")
	}
	return c.JSON(200, "kong")
}
