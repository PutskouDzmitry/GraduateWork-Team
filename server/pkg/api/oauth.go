package api

import (
	"encoding/json"
	"fmt"
	"github.com/PutskouDzmitry/GraduateWork-Team/server/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"unicode/utf8"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback",
		ClientID:     "166169412991-ht2albsna0smdntqnb5v1n2qfpmrqnkp.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	randomState = "random"
)

func (h Handler) homeTest(c *gin.Context) {
	var html = `<html><body><a href="/auth/loginTest"> Google log In</a></body></html>`
	fmt.Fprint(c.Writer, html)
}

func (h Handler) loginTest(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(c.Writer, c.Request, url, http.StatusTemporaryRedirect)
}

func (h Handler) callback(c *gin.Context) {
	if c.Request.FormValue("state") != randomState {
		logrus.Error("error with read value state")
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, c.Request.FormValue("code"))
	if err != nil {
		logrus.Error("error with detect user token: ", err)
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		logrus.Error("error with response: ", err)
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("error with read body", resp)
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	var result ResponseOauth
	err = json.Unmarshal(content, &result)
	if err != nil {
		logrus.Error("error with unmarshal resp.body: ", err)
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	userAccessToken, err := h.createAccessToken(result)
	if err != nil {
		logrus.Error("error with create access token: ", err)
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	userRefreshToken, err := h.createRefreshToken(result)
	if err != nil {
		logrus.Error("error with create refresh token:", err)
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	c.SetCookie("refreshToken", userRefreshToken, 20, "/", "localhost", true, true)
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"id": userAccessToken,
	//})
	redirectResponse(c, "", "http://localhost:3000/api/map/home", userAccessToken)
}

func changeIdToInt(id string) (int, error) {
	n := new(big.Int)
	n, ok := n.SetString(id, 10)
	if !ok {
		return -1, nil
	}
	var userId int
	var err error
	r := []rune(id)
	if utf8.RuneCountInString(id) >= 5 {
		userId, err = strconv.Atoi(string(r[0:5]))
		if err != nil {
			return -1, nil
		}
	}
	return userId, nil
}

func (h Handler) createUser(result ResponseOauth) (model.User, error) {
	userId, err := changeIdToInt(result.Id)
	if err != nil {
		return model.User{}, err
	}

	newUser := model.User{
		Id:       userId,
		Username: result.Email,
		Password: result.Email,
	}
	_, err = h.authService.CreateUser(newUser)
	if err != nil {
		return model.User{}, err
	}
	return newUser, nil
}

func (h Handler) checkUser(result ResponseOauth) (model.User, error) {
	userId, err := changeIdToInt(result.Id)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Id:       userId,
		Username: result.Email,
		Password: result.Email,
	}
	checkUser := h.authService.CheckUser(user)
	if checkUser == false {
		user, err = h.createUser(result)
		if err != nil {
			return model.User{}, err
		}
	}
	return user, nil
}

func (h Handler) createAccessToken(result ResponseOauth) (string, error) {
	user, err := h.checkUser(result)
	if err != nil {
		return "", err
	}
	return h.authService.GenerateTokenAccessToken(user.Id, user.Username, user.Password)
}

func (h Handler) createRefreshToken(result ResponseOauth) (string, error) {
	user, err := h.checkUser(result)
	if err != nil {
		return "", err
	}
	return h.authService.GenerateTokenRefreshToken(user.Id, user.Username, user.Password)
}

type ResponseOauth struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}
