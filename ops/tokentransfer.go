package ops

import (
	"github.com/dfuse-io/eosio-boot/config"
	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/token"
)

func init() {
	Register("token.transfer", &OpTransferToken{})
}

type OpTransferToken struct {
	From     eos.AccountName
	To       eos.AccountName
	Quantity eos.Asset
	Memo     string
}

func (op *OpTransferToken) Actions(c *config.OpConfig) (out []*eos.Action, err error) {
	act := token.NewTransfer(op.From, op.To, op.Quantity, op.Memo)
	return append(out, act), nil
}
