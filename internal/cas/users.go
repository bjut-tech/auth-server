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

func requestUsing(
	ctx context.Context,
	client *http.Client,
	baseUrl string,
	data *url.Values,
) (*http.Response, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		baseUrl+"/v1/users",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil, err
	}

	fakeIp := utils.RandomIPv6()
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("X-Forwarded-For", fakeIp)
	req.Header.Set("X-Real-IP", fakeIp)

	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode == http.StatusForbidden ||
		resp.StatusCode == http.StatusNotFound ||
		resp.StatusCode == http.StatusInternalServerError ||
		resp.StatusCode == http.StatusServiceUnavailable {
		return nil, &ErrConnectionFailed{}
	}

	return resp, nil
}

func GetUser(ctx context.Context, username string, password string) (*UserPrincipal, error) {
	data := url.Values{
		"username": {username},
		"password": {password},
	}

	// Connection fallback mechanism
	resp, err := requestUsing(ctx, client, canonicalBaseUrl, &data)
	if err != nil {
		resp, err = requestUsing(ctx, client, webvpnBaseUrl, &data)
	}
	if err != nil {
		resp, err = requestUsing(ctx, clientPinned, canonicalBaseUrl, &data)
	}
	if err != nil {
		return nil, &ErrConnectionFailed{}
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
	if name, ok := response.Authentication.Principal.Attributes["name"]; ok && len(name) > 0 {
		user.Name = name[0].(string)
	}
	if email, ok := response.Authentication.Principal.Attributes["email"]; ok && len(email) > 0 {
		user.Email = email[0].(string)
	}
	if phone, ok := response.Authentication.Principal.Attributes["phone"]; ok && len(phone) > 0 {
		user.Phone = phone[0].(string)
	}

	return user, nil
}
