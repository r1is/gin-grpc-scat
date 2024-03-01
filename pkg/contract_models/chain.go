package models

// Chain 链
type Chain struct {
	Name   string `json:"name"`
	NameZH string `json:"name_zh"`
}

var ChainUnknown Chain = Chain{
	Name:   "Unknown",
	NameZH: "未知",
}

var ChainETH Chain = Chain{
	Name:   "ETH",
	NameZH: "以太坊",
}

var ChainBSC Chain = Chain{
	Name:   "BSC",
	NameZH: "币安链",
}

func GetChain(name string) Chain {
	switch name {
	case "ETH":
		return ChainETH
	case "BSC":
		return ChainBSC
	default:
		return ChainUnknown
	}
}
