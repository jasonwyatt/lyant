package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/revel/revel"
)

type Defaults struct {
	github_client_id     string
	github_client_secret string
	github_redirect_url  string
	github_oAuth_url     string
	github_api_user      string
	github_login_url     string
}

func GetDefaults() Defaults {
	defaults := Defaults{}
	defaults.github_client_id = "GITHUB CLIENT ID"      //SET YOUR GITHUB CLIENT ID
	defaults.github_client_secret = "GITHUB SECRET KEY" //SET SECRET KEY FOR GITHUB ACCOUNT
	defaults.github_redirect_url = "http://localhost:9000/callback"
	defaults.github_oAuth_url = "https://github.com/login/oauth/access_token"
	defaults.github_api_user = "https://api.github.com/user"
	defaults.github_login_url = "https://github.com/login/oauth/authorize?scope=user:email&client_id="
	return defaults
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	loginUrl := GetDefaults().github_login_url + GetDefaults().github_client_id
	return c.Render(loginUrl)

}

func (c App) GitLogin() revel.Result {
	c.Params.Query = c.Request.URL.Query()
	fmt.Println(c.Params.Query)
	var code string
	var access_token string
	c.Params.Bind(&code, "code")
	if code != "" {
		fmt.Println(code)
	} else {
		c.Params.Bind(&access_token, "access_token")
		fmt.Println(access_token)
	}

	go callForAuthToken(code)
	return c.Render()
}

type TokenBody struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type UserBody struct {
	UserName  string `json:"login"`
	Userid    int    `json:"id"`
	AvatarUrl string `json:"avatar_url"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	UserEmail string `json:"email"`
}

func getUserDetails(authToken string) (*UserBody, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", GetDefaults().github_api_user, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+authToken)
	resp, _ := client.Do(req)
	if err != nil {
		fmt.Println("Something went wrong :" + err.Error())
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	var userBody = new(UserBody)
	error := json.Unmarshal([]byte(bodyString), &userBody)
	if error != nil {
		fmt.Println("whoops:", err)
	}
	return userBody, err

}

func callForAuthToken(code string) {

	form := url.Values{}
	form.Add("client_id", GetDefaults().github_client_id)
	form.Add("client_secret", GetDefaults().github_client_secret)
	form.Add("code", code)
	form.Add("redirect_uri", GetDefaults().github_redirect_url)

	client := &http.Client{}

	req, err := http.NewRequest("POST", GetDefaults().github_oAuth_url, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	req.PostForm = form
	resp, _ := client.Do(req)
	if err != nil {
		fmt.Println("Something went wrong :" + err.Error())
		return
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Println(string(bodyString))
	s, err := parseAuthResponce([]byte(bodyString))
	user, error := getUserDetails(s.AccessToken)
	if error != nil {
		fmt.Println("Error", error)
	}

	fmt.Println(user.UserName + " THIS IS MY USERNAME")
}

func parseAuthResponce(body []byte) (*TokenBody, error) {
	var s = new(TokenBody)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
