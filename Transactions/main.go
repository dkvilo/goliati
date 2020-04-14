package transactions

// TransactionParticipant structure
type TransactionParticipant struct {
	PublicKey []byte `json:"-"`
	Address []byte `json:"address"`
}

// TransactionData structure
type TransactionData struct {
	Amount float64 `json:"amount"`
}

// Transaction structure
type Transaction struct {
	Sender TransactionParticipant `json:"sender"`
	Receiver TransactionParticipant `json:"receiver"`
	Data TransactionData `json:"data"`
	Timestamp string	`json:"timestamp"`
}

