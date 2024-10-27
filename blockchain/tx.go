package blockchain

// txnoutput is a struct with elements value(value of tokens) and pubkey(public key to access the tokens)
type TxOutput struct {
	Value  int
	PubKey string
}

// txnoutputs is a struct with elements id(of the output to spend) out(index of the output to be spend) sig(signature to verify that the public key of the output belongs to you)
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
