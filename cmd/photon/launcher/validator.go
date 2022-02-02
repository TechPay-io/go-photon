package launcher

import (
	"github.com/pkg/errors"
	cli "gopkg.in/urfave/cli.v1"

	"github.com/TechPay-io/sirius-base/inter/idx"

	"github.com/TechPay-io/go-photon/gossip/emitter"
	"github.com/TechPay-io/go-photon/integration/makegenesis"
	"github.com/TechPay-io/go-photon/inter/validatorpk"
)

var validatorIDFlag = cli.UintFlag{
	Name:  "validator.id",
	Usage: "ID of a validator to create events from",
	Value: 0,
}

var validatorPubkeyFlag = cli.StringFlag{
	Name:  "validator.pubkey",
	Usage: "Public key of a validator to create events from",
	Value: "",
}

var validatorPasswordFlag = cli.StringFlag{
	Name:  "validator.password",
	Usage: "Password to unlock validator private key",
	Value: "",
}

// setValidatorID retrieves the validator ID either from the directly specified
// command line flags or from the keystore if CLI indexed.
func setValidator(ctx *cli.Context, cfg *emitter.Config) error {
	// Extract the current validator address, new flag overriding legacy one
	var validatorID idx.ValidatorID
	var validatorPubkey validatorpk.PubKey
	var err error
	if ctx.GlobalIsSet(FakeNetFlag.Name) {
		var num int
		validatorID, num, err = parseFakeGen(ctx.GlobalString(FakeNetFlag.Name))
		if err != nil {
			return err
		}
		validators := makegenesis.GetFakeValidators(num)
		validatorPubkey = validators.Map()[validatorID].PubKey
	}
	if ctx.GlobalIsSet(validatorIDFlag.Name) {
		validatorID = idx.ValidatorID(ctx.GlobalInt(validatorIDFlag.Name))
	}
	if ctx.GlobalIsSet(validatorPubkeyFlag.Name) {
		validatorPubkey, err = validatorpk.FromString(ctx.GlobalString(validatorPubkeyFlag.Name))
		if err != nil {
			return err
		}
	}

	// Convert the validator into an address and configure it
	if validatorID == 0 {
		return nil
	}

	if validatorPubkey.Empty() {
		return errors.New("validator public key is not set")
	}

	cfg.Validator.ID = validatorID
	cfg.Validator.PubKey = validatorPubkey
	return nil
}
