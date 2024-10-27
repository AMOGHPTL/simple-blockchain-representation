package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// transaction struct stores the elements id which is a hash value,inputs which is a slice of type tximputs and outputs which is a slice of type txnoutput
type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

// setid finction is used to set the id of the transaction that is applied to the transaction
// here we create a bytes buffer encoded to store the encoded form
// we encode the tx(transaction) and store the encoded form in the bytes buffer we created named as encoded varibale
// we hash the bytes of the encoded variable using sha256 and store the value in hash variable
// so this basically takes the tx(transasction) and creates a hash of it and assigns it to the the tx.ID
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// coinbasetx is a func that is the coin's base when the genisis block is mined
// we first check if the data string is empty if yes then we assign the data by ourself
// now we create two variable as the struct txinput and txputput and give the required values
// now we create a tx variable as the pointer to the struct transaction with the id value as nil
// we now use func setId on the tx varible which set the id(hash) of the current tx(transaction)
// finally we return the tx(transaction) after the id has been assigned
func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{100, to}

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()

	return &tx
}

func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := chain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error: not enough funds")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, to})

	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}

// func IsCoinbase is used to check if the transaction is coinbase transaction or not
// here we return a boolean value for multiple conditions
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}
