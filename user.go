package bitgo

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
	Email struct {
		ID       string `json:"email"`
		Verified bool   `json:"verified"`
	}
	AllowedCoins []string `json:"allowedCoins"`
}

// UserSession -
type UserSession struct {
	Client  string   `json:"client"`
	UserID  string   `json:"user"`
	Scope   []string `json:"scope"`
	Expires string   `json:"expires"`
	Origin  string   `json:"origin"`
	Unlock  struct {
		Time    string `json:"time"`
		Expires string `json:"expires"`
		TxCount int64  `json:"txCount"`
		TxValue int64  `json:"txValue"`
	} `json:"unlock"`
}

// CurrentUser -
func (c *Config) CurrentUser() (*User, error) {
	userList := &User{}
	c.BaseURL = c.BaseURL + v2 + "user/me"
	err := c.SendHTTPRequest("GET", c.BaseURL, userList)
	if err != nil {
		return nil, err
	}
	return userList, nil
}

// CurrentUserSession -
func (c *Config) CurrentUserSession() (*UserSession, error) {
	userSList := &UserSession{}
	c.BaseURL = c.BaseURL + v2 + "user/session"
	err := c.SendHTTPRequest("GET", c.BaseURL, userSList)
	if err != nil {
		return nil, err
	}
	return userSList, nil
}
