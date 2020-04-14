package block

import (
	"time"

	crypto "github.com/dkvilo/goliati/Crypto"
	global "github.com/dkvilo/goliati/Global"
	transactions "github.com/dkvilo/goliati/Transactions"
)

// Block DataType definition
type Block struct {
	Timestemp    string                     `json:"timesemp,omitempty"`
	Hash         string                     `json:"hash,omitempty"`
	PreviousHash string                     `json:"previousHash,omitempty"`
	Nonce        uint32                     `json:"-"`
	Transactions []transactions.Transaction `json:"transactions"`
}

// New Creates empty block
func (block Block) New(ltsBlock Block) Block {

	block.Nonce = (ltsBlock.Nonce + global.NonceIncrementValue)
	block.Timestemp = time.Now().String()
	block.Hash = crypto.GenerateHash(
		string(block.Nonce),
		block.Timestemp,
	)
	// Genesis block doesnot have any transactions
	block.Transactions = make([]transactions.Transaction, 0, 0)
	return block
}

// CreateGenesis Block
func (block *Block) CreateGenesis() Block {
	// block.Timestemp = time.Now().String()
	block.PreviousHash = ""
	block.Nonce = global.NonceInitialValue
	block.Hash = crypto.GenerateHash(
		string(block.Nonce),
		block.Timestemp,
	)
	return *block
}

// AttachTransaction appends transaction to the block
func (block *Block) AttachTransaction(transaction transactions.Transaction) {
	if transaction.Sender.PublicKey != nil && transaction.Sender.Address != nil &&
		transaction.Receiver.PublicKey != nil && transaction.Receiver.Address != nil {
			block.Transactions = append(block.Transactions, transaction)
	}
}

// AttachTransactions appends multiple transactions to the block
func (block *Block) AttachTransactions(transactions []transactions.Transaction) {
	for _, transaction := range transactions {
		if transaction.Sender.PublicKey != nil && transaction.Sender.Address != nil &&
		transaction.Receiver.PublicKey != nil && transaction.Receiver.Address != nil {
			block.Transactions = append(block.Transactions, transaction)
		}
	}
}
