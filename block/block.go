package block

import (
	"time"
)

type Block struct {
	transactions []Transaction
	prevHash     string
	nonce        string
	timestamp    time.Time
}

type Transaction struct {
	sender    string
	amount    int
	timestamp time.Time
}

func GenerateFirstBlock() {

}
