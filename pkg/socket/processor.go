package socket

import (
	"encoding/json"

	"github.com/centrifugal/centrifuge-go"
	"github.com/galt-tr/tx-filter/pkg/common"
	"github.com/galt-tr/tx-filter/pkg/filters"
	"github.com/galt-tr/tx-filter/pkg/woc"
	"github.com/libsv/go-bt"
)

type Processor struct {
	FilterType common.TransactionType
}

func NewProcessor(txType common.TransactionType) *Processor {
	return &Processor{FilterType: txType}
}

func (p *Processor) FilterMempoolPublishEvent(e centrifuge.ServerPublishEvent) (*bt.Tx, error) {
	transaction := woc.MempoolTransaction{}
	err := json.Unmarshal(e.Data, &transaction)
	if err != nil {
		return nil, err
	}

	switch p.FilterType {
	case common.Metanet:
		return filters.Metanet(&transaction)
	case common.PubKeyHash:
		return filters.PubKeyHash(&transaction)
	case common.RareCandyFrogCartel:
		return filters.RareCandyFrogCartel(&transaction)
	}
	return nil, nil
}
