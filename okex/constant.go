package okex

const (
	// 产品类型
	InstTypeSPOT    = "SPOT"    // 币币
	InstTypeMARGIN  = "MARGIN"  // 币币杠杆
	InstTypeSWAP    = "SWAP"    // 永续合约
	InstTypeFUTURES = "FUTURES" // 交割合约
	InstTypeOPTION  = "OPTION"  // 期权

	// 仓位类型
	MgnModeIsolated = "isolated" // 逐仓
	MgnModeCross    = "cross"    // 全仓

	// 仅交割/永续有效
	CtTypeLinear  = "linear"  // 正向合约
	CtTypeInverse = "inverse" // 反向合约
)
