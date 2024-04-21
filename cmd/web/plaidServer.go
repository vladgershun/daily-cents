package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/plaid/plaid-go/v23/plaid"
)

type PlaidAccount []struct {
	AccountID string `json:"account_id"`
	Balances  struct {
		Available              int     `json:"available"`
		Current                float64 `json:"current"`
		IsoCurrencyCode        string  `json:"iso_currency_code"`
		Limit                  any     `json:"limit"`
		UnofficialCurrencyCode any     `json:"unofficial_currency_code"`
	} `json:"balances"`
	Mask                string `json:"mask"`
	Name                string `json:"name"`
	OfficialName        string `json:"official_name"`
	PersistentAccountID string `json:"persistent_account_id,omitempty"`
	Subtype             string `json:"subtype"`
	Type                string `json:"type"`
}

var (
	PLAID_CLIENT_ID = ""
	PLAID_SECRET    = ""
	PLAID_ENV       = ""
	// PLAID_PRODUCTS                       = ""
	// PLAID_COUNTRY_CODES                  = ""
	// PLAID_REDIRECT_URI                   = ""
	// APP_PORT                             = ""
	client      *plaid.APIClient
	accessToken string
	publicToken string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}

	PLAID_CLIENT_ID = os.Getenv("PLAID_CLIENT_ID")
	PLAID_SECRET = os.Getenv("PLAID_SECRET")

	if PLAID_CLIENT_ID == "" || PLAID_SECRET == "" {
		log.Fatal("Error: PLAID_SECRET or PLAID_CLIENT_ID is not set. Did you copy .env.example to .env and fill it out?")
	}

	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", "66233ad1a0f038001c8a2916")
	configuration.AddDefaultHeader("PLAID-SECRET", "373eebe6c548c9a4928794cffdb4a1")
	configuration.UseEnvironment(plaid.Sandbox)
	client = plaid.NewAPIClient(configuration)

}

func balance(w http.ResponseWriter, r *http.Request) {

	createPublicTokenSandbox()

	ctx := context.Background()

	balancesGetResp, _, err := client.PlaidApi.AccountsBalanceGet(ctx).AccountsBalanceGetRequest(
		*plaid.NewAccountsBalanceGetRequest(accessToken),
	).Execute()

	if err != nil {
		fmt.Println("Could not get balance")
		return
	}

	accounts := balancesGetResp.Accounts

	accountsJSON, err := json.Marshal(accounts)
	if err != nil {
		fmt.Println("Error marshaling accounts:", err)
		return
	}

	var balanceResp PlaidAccount
	err = json.Unmarshal(accountsJSON, &balanceResp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(balanceResp)
}

func createPublicTokenSandbox() {
	ctx := context.Background()

	sandboxPublicTokenResp, _, err := client.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(
		*plaid.NewSandboxPublicTokenCreateRequest(
			"ins_109508",
			[]plaid.Products{plaid.PRODUCTS_TRANSACTIONS},
		),
	).Execute()

	if err != nil {
		plaidError, plaidErr := plaid.ToPlaidError(err)
		if plaidErr == nil {
			fmt.Println(plaidError)
		}
	}

	exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(sandboxPublicTokenResp.GetPublicToken()),
	).Execute()

	accessToken = exchangePublicTokenResp.GetAccessToken()
}

func createLinkToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: "vlad",
	}

	request := plaid.NewLinkTokenCreateRequest(
		"Plaid Quickstart",
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
		user,
	)

	request.SetProducts([]plaid.Products{plaid.PRODUCTS_TRANSACTIONS})
	request.SetRedirectUri("http://localhost:2000")

	linkTokenCreateResp, _, err := client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()
	if err != nil {
		plaidError, plaidErr := plaid.ToPlaidError(err)
		if plaidErr == nil {
			fmt.Println(plaidError)
		}
	}

	linkToken := linkTokenCreateResp.GetLinkToken()
	w.Header().Add("link_token", linkToken)
}
