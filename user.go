package bitgo

import "encoding/json"

// User -
type User struct {
	User UserList `json:"user"`
}

// UserList holds the user details
type UserList struct {
	ID       string `json:"id"`
	IsActive bool   `json:"isActive"`
	Username string `json:"username"`
	Name     struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Full  string `json:"full"`
	} `json:"name"`
}

// CurrentUser -
func (c *Config) CurrentUser() (*User, error) {
	userList := User{}
	c.BaseURL = c.BaseURL + v2

	data, err := c.SendHTTPRequest("GET", c.BaseURL+"user/me")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(data), &userList)
	if err != nil {
		return nil, err
	}
	return &userList, nil
}
