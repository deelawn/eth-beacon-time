package time

import (
	"time"

	"github.com/deelawn/eth-beacon-time/beacon"
)

// ToSlot returns last slot number to begin before or at `slotTime`.
func ToSlot(genesisTime, slotTime uint64) uint64 {

	now := uint64(time.Now().Unix())
	currentSlot := (now - genesisTime) / beacon.SecondsPerSlot
	slotsSinceTime := (now - slotTime) / beacon.SecondsPerSlot

	return currentSlot - slotsSinceTime
}
