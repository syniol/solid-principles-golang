package after

import (
	"encoding/json"
	"io"
	"net/http"
)

type OpenBanking struct {
	httpClient *http.Client
}

func NewOpenBanking() PaymentTransaction[any] {
	return &OpenBanking{
		httpClient: http.DefaultClient,
	}
}

func (s *OpenBanking) TransactionHistory() []any {
	rq, _ := http.NewRequest(
		http.MethodGet,
		"/transactions",
		nil,
	)

	resp, _ := s.httpClient.Do(rq)
	resBody, _ := io.ReadAll(resp.Body)

	var transactions []any
	json.Unmarshal(resBody, &transactions)

	return transactions
}
