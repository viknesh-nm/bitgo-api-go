package bitgo

import (
	"encoding/json"
)

// Wallets holds the list of wallets
type Wallets struct {
	Coin    string          `json:"coin"`
	Wallets []WalletDetails `json:"wallets"`
}

// WalletDetails holds the paricular wallet details
type WalletDetails struct {
	ID                       string           `json:"id"`
	Label                    string           `json:"label"`
	Coin                     string           `json:"coin"`
	M                        int64            `json:"m"`
	N                        int64            `json:"n"`
	Keys                     []string         `json:"keys"`
	Tags                     []string         `json:"tags"`
	Users                    []UserDetails    `json:"users"`
	TransactionNotifications bool             `json:"disableTransactionNotifications"`
	Deleted                  bool             `json:"deleted"`
	ApprovalsRequired        int64            `json:"approvalsRequired"`
	Balance                  string           `json:"balanceString"`
	ConfirmedBalance         string           `json:"confirmedBalanceString"`
	SpendableBalance         string           `json:"spendableBalanceString"`
	CoinSpecific             *json.RawMessage `json:"coinSpecific"`
}

// UserDetails holds the user details
type UserDetails struct {
	UserID     string   `json:"user"`
	Permission []string `json:"permission"`
}

// ListWallets -
func (c *Config) ListWallets(coin string) (*Wallets, error) {
	wallets := &Wallets{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet"
	data, err := c.SendHTTPRequest("GET", c.BaseURL)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(data), &wallets)
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

// GetWallet -
func (c *Config) GetWallet(coin string, walletID string) (*WalletDetails, error) {
	wallet := &WalletDetails{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID
	data, err := c.SendHTTPRequest("GET", c.BaseURL)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(data), &wallet)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}
