package dip

import (
	"encoding/json"
	"io"
	"net/http"
)

type Stripe struct {
	httpClient *http.Client
}

func NewStripe() PaymentTransaction[any] {
	return &Stripe{
		httpClient: http.DefaultClient,
	}
}

func (s *Stripe) TransactionHistory() []any {
	rq, _ := http.NewRequest(
		http.MethodGet,
		"/v1/reporting/transactions",
		nil,
	)

	resp, _ := s.httpClient.Do(rq)
	resBody, _ := io.ReadAll(resp.Body)

	var transactions []any
	json.Unmarshal(resBody, &transactions)

	return transactions
}
