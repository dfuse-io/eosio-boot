package boot

import (
	"context"
	"fmt"
	"strings"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"go.uber.org/zap"
)

func (b *Boot) setKeys() error {
	if b.keyBag == nil {
		zlog.Info("key bag not preset")
		b.keyBag = eos.NewKeyBag()
	}

	for label, privKey := range b.bootseqKeys {
		privKeyStr := privKey.String()
		zlog.Info("adding bootseq key to keybag",
			zap.String("key_tag", label),
			zap.String("pub_key", privKey.PublicKey().String()),
			zap.String("priv_key_prefix", privKey.String()[:4]),
			zap.String("priv_key", privKey.String()[len(privKey.String())-4:]),
		)
		b.keyBag.Add(privKeyStr)
	}

	return nil
}

func (b *Boot) attachKeysOnTargetNode(ctx context.Context) error {
	// Store keys in wallet, to sign `SetCode` and friends..
	b.targetNetAPI.SetSigner(b.keyBag)
	return nil
}

func (b *Boot) parseBootseqKeys() error {
	for label, key := range b.bootSequence.Keys {
		privKey, err := ecc.NewPrivateKey(strings.TrimSpace(key))
		if err != nil {
			return fmt.Errorf("unable to correctly decode %q private key %q: %s", label, key, err)
		}
		b.bootseqKeys[label] = privKey
	}
	return nil
}


func (b *Boot) GetBootseqKey(label string) (*ecc.PrivateKey, error) {
	if _, found := b.bootseqKeys[label]; found {
		return b.bootseqKeys[label], nil
	}
	return nil, fmt.Errorf("bootseq does not contain key with label %q", label)
}
