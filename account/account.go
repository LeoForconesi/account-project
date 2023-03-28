package account

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const ACCOUNT_API_URL = "http://localhost:8080/v1"

type Account interface {
	Fetch(ctx context.Context, accountId string)
	Delete(ctx context.Context, accountId string)
	Create(ctx context.Context)
}

// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

func (a *AccountAttributes) Fetch(ctx context.Context, accountId string) AccountAttributes {
	const URL = ACCOUNT_API_URL + "/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	var account AccountAttributes
	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is ", byteCount)
	fmt.Println("response => ", responseString.String())
	error := json.Unmarshal(content, account)
	if error != nil {
		panic(error)
	}
	return account
}
