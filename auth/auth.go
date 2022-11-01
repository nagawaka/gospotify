package auth

import (
	"encoding/json"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"io"
	"github.com/spf13/viper"
	"spoti/crawl/utils"
	"time"
	"strings"
)

func requestToken(code string) {
	data := url.Values{}
    data.Set("code", code)
    data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", "http://localhost:3333/callback")
	
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout:timeout,
	}

	authStr := fmt.Sprintf("%s:%s", viper.Get("CLIENT_ID"), viper.Get("CLIENT_SECRET"))
	s := fmt.Sprintf("Basic %s", base64.URLEncoding.EncodeToString([]byte(authStr)))
	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", s)

	resp, err := client.Do(request)
	
	if err != nil {
		print(err)
	}

    defer resp.Body.Close()
    var responseData map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&responseData)
    if err != nil {
        print(err)
    }
    fmt.Println(responseData)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	state := utils.RandomString(16)
	viper.Set("CLIENT_STATE", state)
	values := url.Values{}
	values.Add("response_type", "code")
	values.Add("client_id", fmt.Sprintf("%s", viper.Get("CLIENT_ID")))
	values.Add("scope", "user-read-private")
	values.Add("redirect_uri", "http://localhost:3333/callback")
	values.Add("state", state)
	query := values.Encode()
	http.Redirect(w, r, fmt.Sprintf("https://accounts.spotify.com/authorize?%s",query), http.StatusFound)
}

func Callback(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	state := values.Get("state")
	code := values.Get("code")

	if state != viper.Get("CLIENT_STATE") {
		io.WriteString(w, "CLIENT STATE ERROR")
	} else {
		requestToken(code)
	}
}