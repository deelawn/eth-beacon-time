package slot

import (
	"time"

	"github.com/deelawn/eth-beacon-time/beacon"
)

func ToTime(genesisTime, slot uint64) time.Time {

	timeSinceGenesis := slot * beacon.SecondsPerSlot
	timeAtSlot := genesisTime + timeSinceGenesis

	return time.Unix(int64(timeAtSlot), 0)
}
