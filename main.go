package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	block "github.com/dkvilo/goliati/Block"
	blockchain "github.com/dkvilo/goliati/BlockChain"
	p2p "github.com/dkvilo/goliati/P2p"
	transactions "github.com/dkvilo/goliati/Transactions"
)

func main() {

	var chain blockchain.Chain = blockchain.Chain{}

	chain = chain.New()

	nb := block.Block{}
	nb = nb.New(chain.GetLastBlock())

	nb.AttachTransactions([]transactions.Transaction{
		{
			Sender: transactions.TransactionParticipant{
				PublicKey: []byte("david123"),
				Address:   []byte("davids-wallet-address"),
			},
			Receiver: transactions.TransactionParticipant{
				PublicKey: []byte("user123"),
				Address:   []byte("user-wallet-address"),
			},
			Data:      transactions.TransactionData{Amount: 0.1},
			Timestamp: time.Now().String(),
		},
		{
			Sender: transactions.TransactionParticipant{
				PublicKey: []byte("david123"),
				Address:   []byte("davids-wallet-address"),
			},
			Receiver: transactions.TransactionParticipant{
				PublicKey: []byte("user123"),
				Address:   []byte("user-wallet-address"),
			},
			Data:      transactions.TransactionData{Amount: 0.1},
			Timestamp: time.Now().String(),
		},
	})
	chain.AddBlock(nb)

	nb2 := block.Block{}
	nb2 = nb.New(chain.GetLastBlock())
	nb2.AttachTransaction(
		transactions.Transaction{
			Sender:    transactions.TransactionParticipant{PublicKey: []byte("davids-wallet-address")},
			Receiver:  transactions.TransactionParticipant{PublicKey: []byte("receivers-wallet-address")},
			Data:      transactions.TransactionData{Amount: 1},
			Timestamp: time.Now().String(),
		},
	)
	chain.AddBlock(nb2)

	nb3 := block.Block{}
	nb3 = nb.New(chain.GetLastBlock())

	nb3.AttachTransaction(
		transactions.Transaction{
			Sender:    transactions.TransactionParticipant{PublicKey: []byte("davids-wallet-address")},
			Receiver:  transactions.TransactionParticipant{PublicKey: []byte("receivers-wallet-address")},
			Data:      transactions.TransactionData{Amount: 2},
			Timestamp: time.Now().String(),
		},
	)
	chain.AddBlock(nb3)

	// Skip GenesisBlock
	jsonData, err := json.MarshalIndent(chain.Ledger[1:], "", " ")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/transaction", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
				recover()
			}
			fmt.Fprintf(w, string("{\"success\": true }"))
			fmt.Println(string(body))
		}

		if req.Method == "GET" {
			fmt.Fprintf(w, string(jsonData))
		}

	})

	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()

	if (strings.ToLower(*flagMode) == "server") {
		go p2p.Start("server")
		http.ListenAndServe(":8080", nil)
	} else {
		p2p.Start("client")
	}
}
