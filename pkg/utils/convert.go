package utils

import (
	"github.com/galt-tr/tx-filter/pkg/woc"
	"github.com/libsv/go-bt"
)

func ConvertMempoolTxFromWoc(mempoolTx *woc.MempoolTransaction) (*bt.Tx, error) {
	return bt.NewTxFromString(mempoolTx.Hex)
}
