package bitgo

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// Wallets holds the list of wallets
type Wallets struct {
	Coin    string          `json:"coin"`
	Wallets []WalletDetails `json:"wallets"`
}

// WalletDetails holds the paricular wallet details
type WalletDetails struct {
	ID                       string        `json:"id"`
	WalletID                 string        `json:"wallet"`
	Address                  string        `json:"address"`
	Label                    string        `json:"label"`
	Coin                     string        `json:"coin"`
	M                        int64         `json:"m"`
	N                        int64         `json:"n"`
	Keys                     []string      `json:"keys"`
	Tags                     []string      `json:"tags"`
	Users                    []UserDetails `json:"users"`
	TransactionNotifications bool          `json:"disableTransactionNotifications"`
	Deleted                  bool          `json:"deleted"`
	ApprovalsRequired        int64         `json:"approvalsRequired"`
	Balance                  string        `json:"balanceString"`
	ConfirmedBalance         string        `json:"confirmedBalanceString"`
	SpendableBalance         string        `json:"spendableBalanceString"`
	CoinSpecific             struct {
		RedeemScript string `json:"redeemScript"`
	} `json:"coinSpecific"`
}

// WalletAddress -
type WalletAddress struct {
	Limit               int64           `json:"limit"`
	Coin                string          `json:"coin"`
	Count               int64           `json:"count"`
	PendingAddressCount int64           `json:"pendingAddressCount"`
	TotalAddressCount   int64           `json:"totalAddressCount"`
	Addresses           []WalletDetails `json:"addresses"`
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
	TransID       string         `json:"txid"`
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

// TransHistory -
type TransHistory struct {
	Date   string `json:"date"`
	Action string `json:"action"`
}

// TransEntry -
type TransEntry struct {
	Address     string `json:"address"`
	Wallet      string `json:"wallet"`
	Value       int64  `json:"value"`
	ValueString string `json:"valueString"`
	Inputs      int64  `json:"inputs"`
	Outputs     int64  `json:"outputs"`
}

// TransIO -
type TransIO struct {
	ID          string `json:"id"`
	Address     string `json:"address"`
	Value       int64  `json:"value"`
	ValueString string `json:"valueString"`
	Wallet      string `json:"wallet"`
	Chain       int64  `json:"chain"`
	Index       int64  `json:"index"`
}

// WalletTransactions -
type WalletTransactions struct {
	Transactions []Transaction `json:"transactions"`
}

// Transaction -
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

// Balances -
type Balances struct {
	Balance                int64  `json:"balance"`
	ConfirmedBalance       int64  `json:"confirmedBalance"`
	SpendableBalance       int64  `json:"spendableBalance"`
	BalanceString          string `json:"balanceString"`
	ConfirmedBalanceString string `json:"confirmedBalanceString"`
	SpendableBalanceString string `json:"spendableBalanceString"`
}

// WalletUnspents -
type WalletUnspents struct {
	Coin            string   `json:"coin"`
	NextBatchPrevID string   `json:"nextBatchPrevId"`
	Unspents        []string `json:"unspents"`
}

// Unspent -
type Unspent struct {
	ID           string `json:"id"`
	Address      string `json:"address"`
	Value        int64  `json:"value"`
	BlockHeight  int64  `json:"blockHeight"`
	Date         string `json:"date"`
	Wallet       string `json:"wallet"`
	FromWallet   string `json:"fromWallet"`
	Chain        int64  `json:"chain"`
	Index        int64  `json:"index"`
	RedeemScript string `json:"redeemScript"`
	IsSegwit     bool   `json:"isSegwit"`
}

// MaximumSpendable -
type MaximumSpendable struct {
	MaxSpendable string `json:"maximumSpendable"`
	Coin         string `json:"coin"`
}

// WalletShare -
type WalletShare struct {
	Incoming []ShareDetails `json:"incoming"`
	Outgoing []ShareDetails `json:"outgoing"`
}

// ShareDetails -
type ShareDetails struct {
	ID          string `json:"id"`
	Coin        string `json:"coin"`
	Wallet      string `json:"wallet"`
	WalletLabel string `json:"walletLabel"`
	FromUser    string `json:"fromUser"`
	ToUser      string `json:"toUser"`
	Permissions string `json:"permissions"`
	State       string `json:"state"`
}

// UserWebhooks -
type UserWebhooks struct {
	Webhooks []UWebHook `json:"webhooks"`
}

// UWebHook -
type UWebHook struct {
	ID                       string `json:"id"`
	Label                    string `json:"label"`
	Created                  string `json:"created"`
	Coin                     string `json:"coin"`
	Type                     string `json:"type"`
	URL                      string `json:"url"`
	State                    string `json:"state"`
	NumConfirmations         int64  `json:"numConfirmations"`
	Version                  int64  `json:"version"`
	SuccessiveFailedAttempts int64  `json:"successiveFailedAttempts"`
}

// GenerateWalletRequest -
type GenerateWalletRequest struct {
	Label                           string `json:"label,mandatory"`
	Passphrase                      string `json:"passphrase,mandatory"`
	UserKey                         string `json:"userKey"`
	BackupXpub                      string `json:"backupXpub"`
	BackupXpubProvider              string `json:"backupXpubProvider"`
	Enterprise                      string `json:"enterprise"`
	DisableTransactionNotifications bool   `json:"disableTransactionNotifications"`
	GasPrice                        int64  `json:"gasPrice"`
	PasscodeEncryptionCod           string `json:"passcodeEncryptionCod"`
}

// GResponse -
type GResponse struct {
	Wallet struct {
		WalletList GWalletResponse `json:"_wallet"`
	} `json:"wallet"`
}

// GWalletResponse -
type GWalletResponse struct {
	ID    string              `json:"id"`
	Label string              `json:"label"`
	Coin  string              `json:"coin"`
	Users []WalletUserDetails `json:"users"`
	Keys  []string            `json:"keys"`
	Balances
}

// WalletUserDetails -
type WalletUserDetails struct {
	User        string   `json:"user"`
	Permissions []string `json:"permissions"`
}

// AddWalletRequest -
type AddWalletRequest struct {
	Label                          string   `json:"label,mandatory"`
	M                              int64    `json:"m,mandatory"`
	N                              int64    `json:"n,mandatory"`
	Keys                           []string `json:"keys,mandatory"`
	Enterprise                     string   `json:"enterprise"`
	IsCold                         bool     `json:"isCold"`
	DisableTransactionNotification bool     `json:"disableTransactionNotification"`
}

// GenerateWallet -
func (c *Config) GenerateWallet(coin string, bRequest GenerateWalletRequest) (*GWalletResponse, error) {

	gWallets := GWalletResponse{}
	walletDetails := &GResponse{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet/generate"

	bodyRequest, err := json.Marshal(bRequest)
	if err != nil {
		return nil, err
	}
	c.body = bodyRequest

	err = c.SendHTTPRequest("POST", c.BaseURL, walletDetails)
	if err != nil {
		return nil, err
	}
	gWallets = walletDetails.Wallet.WalletList
	return &gWallets, nil
}

// AddWallet -
func (c *Config) AddWallet(coin string, bRequest AddWalletRequest) (*GWalletResponse, error) {

	gWallets := GWalletResponse{}
	walletDetails := &GResponse{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet"

	bodyRequest, err := json.Marshal(bRequest)
	if err != nil {
		return nil, err
	}
	c.body = bodyRequest

	err = c.SendHTTPRequest("POST", c.BaseURL, walletDetails)
	if err != nil {
		return nil, err
	}
	gWallets = walletDetails.Wallet.WalletList

	return &gWallets, nil
}

// ListWallets -
func (c *Config) ListWallets(coin string, params ...QParams) (*Wallets, error) {
	wallets := &Wallets{}
	p := url.Values{}
	if params != nil {
		p.Set("allTokens", strconv.FormatBool(params[0].AllTokens))
		p.Set("limit", strconv.Itoa(params[0].Limit))
		p.Set("prevId", params[0].PrevID)
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet?" + p.Encode()
	} else {
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet"
	}
	err := c.SendHTTPRequest("GET", c.BaseURL, wallets)
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

// GetWalletByID -
func (c *Config) GetWalletByID(coin, walletID string, params ...QParams) (*WalletDetails, error) {
	wallet := &WalletDetails{}
	p := url.Values{}
	if params != nil {
		p.Set("allTokens", strconv.FormatBool(params[0].AllTokens))
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "?" + p.Encode()
	} else {
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID
	}
	err := c.SendHTTPRequest("GET", c.BaseURL, wallet)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

// GetWalletByAddress -
func (c *Config) GetWalletByAddress(coin, walletAddress string) (*WalletDetails, error) {
	wallet := &WalletDetails{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet/address/" + walletAddress
	err := c.SendHTTPRequest("GET", c.BaseURL, wallet)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

// WalletWebhooks -
func (c *Config) WalletWebhooks(coin, walletID string, params ...QParams) (*WebHookWallets, error) {
	webHookWallets := &WebHookWallets{}
	p := url.Values{}
	if params != nil {
		p.Set("allTokens", strconv.FormatBool(params[0].AllTokens))
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/webhooks?" + p.Encode()
	} else {
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/webhooks"
	}
	err := c.SendHTTPRequest("GET", c.BaseURL, webHookWallets)
	if err != nil {
		return nil, err
	}
	return webHookWallets, nil
}

// ListWalletTransfer -
func (c *Config) ListWalletTransfer(coin, walletID string, params ...QParams) (*WalletTransfer, error) {
	walletTransfer := &WalletTransfer{}
	p := url.Values{}
	if params != nil {
		p.Set("prevId", params[0].PrevID)
		p.Set("allTokens", strconv.FormatBool(params[0].AllTokens))
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/transfer?" + p.Encode()
	} else {
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/transfer"
	}
	err := c.SendHTTPRequest("GET", c.BaseURL, walletTransfer)
	if err != nil {
		return nil, err
	}
	return walletTransfer, nil
}

// ListWalletUnspents -
func (c *Config) ListWalletUnspents(coin, walletID string, params ...QParams) (*WalletUnspents, error) {
	walletUnspents := &WalletUnspents{}
	p := url.Values{}
	if params != nil {
		p.Set("prevId", params[0].PrevID)
		p.Set("minValue", strconv.Itoa(params[0].MinValue))
		p.Set("maxValue", strconv.Itoa(params[0].MaxValue))
		p.Set("minHeight", strconv.Itoa(params[0].MinHeight))
		p.Set("minConfirms", strconv.Itoa(params[0].MinConfirms))
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/unspents?" + p.Encode()
	} else {
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/unspents"
	}
	err := c.SendHTTPRequest("GET", c.BaseURL, walletUnspents)
	if err != nil {
		return nil, err
	}
	return walletUnspents, nil
}

// GetWalletTransferByID -
func (c *Config) GetWalletTransferByID(coin, walletID, transferID string) (*Transfer, error) {
	walletTransfer := &Transfer{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/transfer/" + transferID
	err := c.SendHTTPRequest("GET", c.BaseURL, walletTransfer)
	if err != nil {
		return nil, err
	}
	return walletTransfer, nil
}

// GetWalletTransferBySequenceID -
func (c *Config) GetWalletTransferBySequenceID(coin, walletID, sequenceID string) (*Transfer, error) {
	walletTransfer := &Transfer{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/transfer/sequenceId/" + sequenceID
	err := c.SendHTTPRequest("GET", c.BaseURL, walletTransfer)
	if err != nil {
		return nil, err
	}
	return walletTransfer, nil
}

// ListWalletTransactions -
func (c *Config) ListWalletTransactions(coin, walletID string, params ...QParams) (*WalletTransactions, error) {
	walletTransactions := &WalletTransactions{}
	p := url.Values{}
	if params != nil {
		p.Set("allTokens", strconv.FormatBool(params[0].AllTokens))
		p.Set("prevId", params[0].PrevID)
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/tx?" + p.Encode()
	} else {
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/tx"
	}
	err := c.SendHTTPRequest("GET", c.BaseURL, walletTransactions)
	if err != nil {
		return nil, err
	}
	return walletTransactions, nil
}

// GetWalletTransactionByID -
func (c *Config) GetWalletTransactionByID(coin, walletID, transID string) (*Transaction, error) {
	walletTransactions := &Transaction{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/tx/" + transID
	err := c.SendHTTPRequest("GET", c.BaseURL, walletTransactions)
	if err != nil {
		return nil, err
	}
	return walletTransactions, nil
}

// GetTotalBalances -
func (c *Config) GetTotalBalances(coin string) (*Balances, error) {
	balances := &Balances{}
	c.BaseURL = c.BaseURL + v2 + coin + "/wallet/balances"
	err := c.SendHTTPRequest("GET", c.BaseURL, balances)
	if err != nil {
		return nil, err
	}
	return balances, nil
}

// GetMaximumSpendable -
func (c *Config) GetMaximumSpendable(coin, walletID string, params ...QParams) (*MaximumSpendable, error) {
	max := &MaximumSpendable{}
	p := url.Values{}
	if params != nil {
		p.Set("prevId", params[0].PrevID)
		p.Set("limit", strconv.Itoa(params[0].Limit))
		p.Set("minValue", strconv.Itoa(params[0].MinValue))
		p.Set("maxValue", strconv.Itoa(params[0].MaxValue))
		p.Set("minHeight", strconv.Itoa(params[0].MinHeight))
		p.Set("minConfirms", strconv.Itoa(params[0].MinConfirms))
		p.Set("feeRate", strconv.Itoa(params[0].FeeRate))
		p.Set("enforceMinConfirmsForChange", strconv.FormatBool(params[0].EnforceMinConfirms))
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/maximumSpendable?" + p.Encode()
	} else {
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/maximumSpendable"
	}
	err := c.SendHTTPRequest("GET", c.BaseURL, max)
	if err != nil {
		return nil, err
	}
	return max, nil
}

// ListWalletAddresses -
func (c *Config) ListWalletAddresses(coin, walletID string, params ...QParams) (*WalletAddress, error) {
	addresses := &WalletAddress{}
	p := url.Values{}
	if params != nil {
		p.Set("prevId", params[0].PrevID)
		p.Set("limit", strconv.Itoa(params[0].Limit))
		p.Set("sortOrder", strconv.Itoa(params[0].SortOrder)) // sort order should be either -1 or 1
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/addresses?" + p.Encode()
	} else {
		c.BaseURL = c.BaseURL + v2 + coin + "/wallet/" + walletID + "/addresses"
	}
	err := c.SendHTTPRequest("GET", c.BaseURL, addresses)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

// ListWalletShares -
func (c *Config) ListWalletShares(coin string) (*WalletShare, error) {
	walletShare := &WalletShare{}
	c.BaseURL = c.BaseURL + v2 + coin + "/walletshare"
	err := c.SendHTTPRequest("GET", c.BaseURL, walletShare)
	if err != nil {
		return nil, err
	}
	return walletShare, nil
}

// GetWalletShareByID -
func (c *Config) GetWalletShareByID(coin, shareID string) (*ShareDetails, error) {
	walletShare := &ShareDetails{}
	c.BaseURL = c.BaseURL + v2 + coin + "/walletshare/" + shareID
	err := c.SendHTTPRequest("GET", c.BaseURL, walletShare)
	if err != nil {
		return nil, err
	}
	return walletShare, nil
}

// ListUserWebhooks -
func (c *Config) ListUserWebhooks(coin string) (*UserWebhooks, error) {
	userWebhooks := &UserWebhooks{}
	c.BaseURL = c.BaseURL + v2 + coin + "/webhooks"
	err := c.SendHTTPRequest("GET", c.BaseURL, userWebhooks)
	if err != nil {
		return nil, err
	}
	return userWebhooks, nil
}
