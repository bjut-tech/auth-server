package cas

import (
	"context"
	"encoding/json"
	"github.com/bjut-tech/auth-server/internal/utils"
	"net/http"
	"net/url"
	"strings"
)

type UserPrincipal struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type getUserResponse struct {
	Authentication struct {
		Principal struct {
			Id         string                   `json:"id"`
			Attributes map[string][]interface{} `json:"attributes"`
		} `json:"principal"`
	} `json:"authentication"`
}

func GetUser(ctx context.Context, username string, password string) (*UserPrincipal, error) {
	data := url.Values{
		"username": {username},
		"password": {password},
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://cas.bjut.edu.cn/v1/users",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("X-Forwarded-For", utils.RandomIPv6())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, &ErrConnectionFailed{err}
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, &ErrInvalidCredentials{}
	} else if resp.StatusCode == http.StatusLocked {
		return nil, &ErrThrottled{}
	} else if resp.StatusCode != http.StatusOK {
		return nil, &ErrUnknown{StatusCode: resp.StatusCode}
	}

	defer resp.Body.Close()
	var response getUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, &ErrUnexpectedResponse{err}
	}

	user := &UserPrincipal{
		Id: response.Authentication.Principal.Id,
	}
	if name, ok := response.Authentication.Principal.Attributes["name"]; ok {
		user.Name = name[0].(string)
	}
	if email, ok := response.Authentication.Principal.Attributes["email"]; ok {
		user.Email = email[0].(string)
	}
	if phone, ok := response.Authentication.Principal.Attributes["phone"]; ok {
		user.Phone = phone[0].(string)
	}

	return user, nil
}
