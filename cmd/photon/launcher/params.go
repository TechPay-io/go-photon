package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://6ae039d34e69dd61fdafcc40198e1cab264afd4e17d22f80b19f7d4d3b67ea36de11ee9b5702b7122d113ec135494c063d68b5b2156b4730fbcf4aaa6581b7d5@178.62.69.60:5050",
		"enode://c8cd5a794e64b4d1c08fb38b35630e981805c959921cec4ff3edd621e105babbe57d4ad2687884a3a00c02bba99d54ab7cebfb519dfc1fe6b0095ecae7971192@167.99.199.7:5050",
		"enode://8ca6e8b391eefe43c31b23368d85ac886030ffebc4cfa4167bc0bb2e818a4136f9ade8a328bd3ffdf3a2dec978cfb6222484c478d9125d9cbddfbb548b2576fa@142.93.33.192:5050",
		"enode://a4b35611552fbb0506c35d4c09b975a2f9a5d50a0dc23f7dde879afe5f86756a6e4cc223837654317b258eec5f8712f2b7da1b6dcce4c2ae9c1ea6c3e7df11fd@159.65.52.165:5050",
		"enode://fc0eac29900b92213d4d6882566ba96668fd54bcce81153c45183d92c372e6152d9eebe0928c9f5968327885fb8b16e2c2d7cbe25c9bdc5c9eea96b292298f31@46.101.28.195:5050",
		"enode://0b22327fc7b0d63a25b5f16389f12f84b81d2de4d3da859d6072621a155d79edfa67943270d1f381a26dc66548257dc10bd5abfcc1f6edca982b6ace9c831e5a@142.93.42.96:5050",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
