package filters

import (
	"strings"

	"github.com/galt-tr/tx-filter/pkg/utils"
	"github.com/galt-tr/tx-filter/pkg/woc"
	"github.com/libsv/go-bt"
)

const RareCandyFrogCartelScriptTemplate = "a914179b4c7a45646a509473df5a444b6e18b723bd148876"

func RareCandyFrogCartel(tx *woc.MempoolTransaction) (*bt.Tx, error) {
	// Loop through all of the outputs and check for pubkeyhash output
	for _, out := range tx.Outputs {
		// if any output contains a pubkeyhash output, include this tx in the filter
		if strings.HasPrefix(out.ScriptPubKey.Hex, RareCandyFrogCartelScriptTemplate) {
			return utils.ConvertMempoolTxFromWoc(tx)
		}
	}
	return nil, nil
}
