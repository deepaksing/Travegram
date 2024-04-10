package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/deepaksing/Travegram/server/api/auth"
	"github.com/deepaksing/Travegram/store"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *ApiV1Service) RegisterUserRoute(g *echo.Group) {
	g.POST("/register", s.RegisterUser)
	g.POST("/singin", s.SiginIn)
}

type SignUp struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SiginIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *ApiV1Service) SiginIn(c echo.Context) error {
	ctx := c.Request().Context()

	signin := &SiginIn{}

	if err := json.NewDecoder(c.Request().Body).Decode(signin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Malformed siginin request")
	}

	user, err := s.store.GetUser(ctx, &store.FindUser{
		Username: &signin.Username,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Incorrect login credentials, please try again")
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect login credentials, please try again")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(signin.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect login credentials, please try again")
	}

	//if user selects remember then expire after the access token duration time else just keep the expiration to time.time
	accessTokenExpirationTime := time.Now().Add(auth.AccessTokenDuration)

	accessToken, err := auth.GenerateAccessToken(user.Username, user.ID, accessTokenExpirationTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to generate tokens, err: %s", err)).SetInternal(err)
	}
	//create cookie
	cookieExp := time.Now().Add(auth.CookieExpDuration)
	setTokenCookie(c, auth.AccessTokenCookieName, accessToken, cookieExp)

	//return json reponse
	return c.JSON(http.StatusOK, user)

}

func (s *ApiV1Service) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	signup := &SignUp{}

	if err := json.NewDecoder(c.Request().Body).Decode(signup); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Malformatted signup request").SetInternal(err)
	}

	userCreate := &store.User{
		Username: signup.Username,
		Name:     signup.Name,
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.DefaultCost)
	fmt.Println(passwordHash)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate password hash").SetInternal(err)
	}

	userCreate.PasswordHash = string(passwordHash)

	user, err := s.store.CreateUser(ctx, userCreate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user").SetInternal(err)
	}

	//get access jwt token
	accessTokenExpirationTime := time.Now().Add(auth.AccessTokenDuration)
	accessToken, err := auth.GenerateAccessToken(user.Username, user.ID, accessTokenExpirationTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to generate tokens, err: %s", err)).SetInternal(err)
	}
	//create cookie
	cookieExp := time.Now().Add(auth.CookieExpDuration)
	setTokenCookie(c, auth.AccessTokenCookieName, accessToken, cookieExp)

	//return json reponse
	return c.JSON(http.StatusOK, user)
}

func setTokenCookie(c echo.Context, name, token string, expiration time.Time) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteStrictMode
	c.SetCookie(cookie)
}
