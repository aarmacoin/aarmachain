package common

const (
	FoundationSmartContract   	= "0x0000000000000000000000000000000000000021"
	SystemSmartContract   		= "0x0000000000000000000000000000000000000051"
	ValidatorsContractAddress 	= "0x0000000000000000000000000000000000000251"

	BlockRewardToValidator 		= 60
	BlockRewardToSystem    		= 20
	BlockRewardToFoundation		= 20
)

var BlackList = map[Address]bool{}
