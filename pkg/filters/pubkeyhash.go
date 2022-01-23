package filters

import (
	"log"

	"github.com/galt-tr/tx-filter/pkg/common"
	"github.com/galt-tr/tx-filter/pkg/utils"
	"github.com/galt-tr/tx-filter/pkg/woc"
	"github.com/libsv/go-bt"
)

func PubKeyHash(tx *woc.MempoolTransaction) (*bt.Tx, error) {
	log.Printf("Attempting to filter for pubkeyhash: %#v", tx)
	// Loop through all of the outputs and check for pubkeyhash output
	for _, out := range tx.Outputs {
		// if any output contains a pubkeyhash output, include this tx in the filter
		if out.ScriptPubKey.Type == string(common.PubKeyHash) {
			return utils.ConvertMempoolTxFromWoc(tx)
		}
	}
	return nil, nil
}
