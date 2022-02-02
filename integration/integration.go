package integration

import (
	"time"

	"github.com/TechPay-io/sirius-base/abft"
	"github.com/TechPay-io/sirius-base/utils/cachescale"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p/simulations/adapters"
	"github.com/status-im/keycard-go/hexutils"

	"github.com/TechPay-io/go-photon/gossip"
	"github.com/TechPay-io/go-photon/inter/validatorpk"
	"github.com/TechPay-io/go-photon/valkeystore"
	"github.com/TechPay-io/go-photon/vecmt"
)

var (
	FlushIDKey = hexutils.HexToBytes("0068c2927bf842c3e9e2f1364494a33a752db334b9a819534bc9f17d2c3b4e5970008ff519d35a86f29fcaa5aae706b75dee871f65f174fcea1747f2915fc92158f6bfbf5eb79f65d16225738594bffb0c")
)

// NewIntegration creates gossip service for the integration test
func NewIntegration(ctx *adapters.ServiceContext, genesis InputGenesis, stack *node.Node) *gossip.Service {
	gossipCfg := gossip.FakeConfig(1, cachescale.Identity)
	cfg := Configs{
		Photon:      gossipCfg,
		PhotonStore: gossip.DefaultStoreConfig(cachescale.Identity),
		Sirius:      abft.DefaultConfig(),
		SiriusStore: abft.DefaultStoreConfig(cachescale.Identity),
		VectorClock: vecmt.DefaultConfig(cachescale.Identity),
	}

	engine, dagIndex, gdb, _, _, blockProc := MakeEngine(DBProducer(ctx.Config.DataDir, cachescale.Identity), genesis, cfg)
	_ = genesis.Close()

	valKeystore := valkeystore.NewDefaultMemKeystore()

	pubKey := validatorpk.PubKey{
		Raw:  crypto.FromECDSAPub(&ctx.Config.PrivateKey.PublicKey),
		Type: validatorpk.Types.Secp256k1,
	}

	// unlock the key
	_ = valKeystore.Add(pubKey, crypto.FromECDSA(ctx.Config.PrivateKey), validatorpk.FakePassword)
	_ = valKeystore.Unlock(pubKey, validatorpk.FakePassword)
	signer := valkeystore.NewSigner(valKeystore)

	// find a genesis validator which corresponds to the key
	for id, v := range gdb.GetEpochState().ValidatorProfiles {
		if v.PubKey.String() == pubKey.String() {
			gossipCfg.Emitter.Validator.ID = id
			gossipCfg.Emitter.Validator.PubKey = v.PubKey
		}
	}

	gossipCfg.Emitter.EmitIntervals.Max = 3 * time.Second
	gossipCfg.Emitter.EmitIntervals.DoublesignProtection = 0

	svc, err := gossip.NewService(stack, gossipCfg, gdb, signer, blockProc, engine, dagIndex)
	if err != nil {
		panic(err)
	}
	err = engine.Bootstrap(svc.GetConsensusCallbacks())
	if err != nil {
		return nil
	}

	return svc
}
