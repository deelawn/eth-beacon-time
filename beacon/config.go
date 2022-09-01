package beacon

const (
	ChainIDMainnet string = "mainnet"
	ChainIDPrater  string = "prater"

	GenesisTimeMainnet uint64 = 1606824000
	GenesisTimePrater  uint64 = 1614588812

	SecondsPerSlot uint64 = 12
)

var GenesisTimes = map[string]uint64{
	ChainIDMainnet: GenesisTimeMainnet,
	ChainIDPrater:  GenesisTimePrater,
}
