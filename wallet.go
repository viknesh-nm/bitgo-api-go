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

// WebHookWallets holds all webhook  wallet details
type WebHookWallets struct {
	WebHooks []WebHookDetails `json:"webhooks"`
}

// WebHookDetails holds particular webhook detail
type WebHookDetails struct {
	ID        string `json:"id"`
	Label     string `json:"label"`
	CreatedAt string `json:"created"`
	WalletID  string `json:"walletId"`
	Coin      string `json:"coin"`
	Type      string `json:"type"`
	URL       string `json:"url"`
}

// WalletTransfer holds the transfer details
type WalletTransfer struct {
	Count     int64      `json:"count"`
	Coin      string     `json:"coin"`
	Transfers []Transfer `json:"transfers"`
}

// Transfer holds for particular transfer details
type Transfer struct {
	ID            string         `json:"id"`
	Coin          string         `json:"coin"`
	WalletID      string         `json:"wallet"`
	TransferID    string         `json:"txid"`
	Height        int64          `json:"height"`
	Date          string         `json:"date"`
	Confirmations int64          `json:"confirmations"`
	Value         int64          `json:"value"`
	BitgoFee      float64        `json:"bitgoFee"`
	USD           float64        `json:"usd"`
	USDRate       float64        `json:"usdRate"`
	State         string         `json:"state"`
	SequenceID    string         `json:"sequenceId"`
	VSize         int64          `json:"vSize"`
	NSegwitInputs int64          `json:"nSegwitInputs"`
	Tags          []string       `json:"tags"`
	ConfirmedTime string         `json:"confirmedTime"`
	CreatedTime   string         `json:"createdTime"`
	History       []TransHistory `json:"history"`
	Entries       []TransEntry   `json:"entries"`
	Inputs        []TransIO      `json:"inputs"`
	Outputs       []TransIO      `json:"outputs"`
}

type TransHistory struct {
	Date   string `json:"date"`
	Action string `json:"action"`
}

type TransEntry struct {
	Address     string `json:"address"`
	Wallet      string `json:"wallet"`
	Value       int64  `json:"value"`
	ValueString string `json:"valueString"`
	Inputs      int64  `json:"inputs"`
	Outputs     int64  `json:"outputs"`
}

type TransIO struct {
	ID          string `json:"id"`
	Address     string `json:"address"`
	Value       int64  `json:"value"`
	ValueString string `json:"valueString"`
	Wallet      string `json:"wallet"`
	Chain       int64  `json:"chain"`
	Index       int64  `json:"index"`
}

type WalletTransactions struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ID               string       `json:"id"`
	NormalizedTxHash string       `json:"normalizedTxHash"`
	Date             string       `json:"date"`
	Hex              string       `json:"hex"`
	BlockHash        string       `json:"blockHash"`
	BlockHeight      int64        `json:"blockHeight"`
	BlockPosition    int64        `json:"blockPosition"`
	Confirmations    int64        `json:"confirmations"`
	Fee              int64        `json:"fee"`
	FeeString        string       `json:"feeString"`
	Size             int64        `json:"size"`
	FromWallet       string       `json:"fromWallet"`
	InputID          []string     `json:"inputIds"`
	Entries          []TransEntry `json:"entries"`
	Inputs           []TransIO    `json:"inputs"`
	Outputs          []TransIO    `json:"outputs"`
}

// ListWallets -
func (c *Config) ListWallets(coin string) (*Wallets, error) {
	wallets := &Wallets{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet"
	err := c.SendHTTPRequest("GET", c.BaseURL, wallets)
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

// GetSingleWallet -
func (c *Config) GetSingleWallet(coin string, walletID string) (*WalletDetails, error) {
	wallet := &WalletDetails{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID
	err := c.SendHTTPRequest("GET", c.BaseURL, wallet)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

// WalletWebhooks -
func (c *Config) WalletWebhooks(coin string, walletID string) (*WebHookWallets, error) {
	webHookWallets := &WebHookWallets{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/webhooks"
	err := c.SendHTTPRequest("GET", c.BaseURL, webHookWallets)
	if err != nil {
		return nil, err
	}
	return webHookWallets, nil
}

// ListWalletTransfer -
func (c *Config) ListWalletTransfer(coin string, walletID string) (*WalletTransfer, error) {
	walletTransfer := &WalletTransfer{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/transfer"
	err := c.SendHTTPRequest("GET", c.BaseURL, walletTransfer)
	if err != nil {
		return nil, err
	}
	return walletTransfer, nil
}

// ListWalletTransactions -
func (c *Config) ListWalletTransactions(coin string, walletID string) (*WalletTransactions, error) {
	walletTransactions := &WalletTransactions{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/tx"
	err := c.SendHTTPRequest("GET", c.BaseURL, walletTransactions)
	if err != nil {
		return nil, err
	}
	return walletTransactions, nil
}
