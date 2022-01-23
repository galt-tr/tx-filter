package filters

import (
	"strings"

	"github.com/galt-tr/tx-filter/pkg/utils"
	"github.com/galt-tr/tx-filter/pkg/woc"
	"github.com/libsv/go-bt"
)

const MetanetScriptTemplate = "14cb030491157b26a570b6ee91e5b068d99c3b72f6"

func Metanet(tx *woc.MempoolTransaction) (*bt.Tx, error) {
	// Loop through all of the outputs and check for pubkeyhash output
	for _, out := range tx.Outputs {
		// if any output contains a pubkeyhash output, include this tx in the filter
		if strings.HasPrefix(out.ScriptPubKey.Hex, MetanetScriptTemplate) {
			return utils.ConvertMempoolTxFromWoc(tx)
		}
	}
	return nil, nil
}
