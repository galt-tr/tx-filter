package woc

type MempoolTransaction struct {
	Hex      string   `json:"hex"`
	Txid     string   `json:"txid"`
	Hash     string   `json:"hash"`
	Size     int      `json:"size"`
	Version  int      `json:"version"`
	LockTime int      `json:"locktime"`
	Inputs   []Input  `json:"vin"`
	Outputs  []Output `json:"vout"`
}

type Input struct {
	Txid      string    `json:"txid"`
	Vout      int       `json:"vout"`
	ScriptSig ScriptSig `json:"scriptSig"`
	Sequence  int       `json:"sequencea"`
}

type Output struct {
	Value        float64      `json:"value"`
	Num          int          `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type ScriptPubKey struct {
	Asm                string   `json:"asm"`
	Hex                string   `json:"hex"`
	RequiredSignatures int      `json:"reqSigs"`
	Type               string   `json:"type"`
	Addresses          []string `json:"addresses"`
	IsTruncated        bool     `json:"isTruncated"`
}
