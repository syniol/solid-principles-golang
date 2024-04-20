package after

import (
	"encoding/json"
	"io"
	"net/http"
)

type PayPal[TransactionType any] struct {
	httpClient *http.Client
}

func NewPayPal[TransactionType any]() PaymentTransaction[TransactionType] {
	return &PayPal[TransactionType]{
		httpClient: http.DefaultClient,
	}
}

func (s *PayPal[TransactionType]) TransactionHistory() []TransactionType {
	rq, _ := http.NewRequest(
		http.MethodGet,
		"/v1/issuing/transactions",
		nil,
	)

	resp, _ := s.httpClient.Do(rq)
	resBody, _ := io.ReadAll(resp.Body)

	var transactions []TransactionType
	json.Unmarshal(resBody, &transactions)

	return transactions
}
