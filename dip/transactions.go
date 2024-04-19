package dip

type PaymentTransaction[T any] interface {
	// TransactionHistory lists all transactions history for a customer
	TransactionHistory() []T
}
