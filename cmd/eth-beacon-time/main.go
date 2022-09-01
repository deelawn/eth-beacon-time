package main

import (
	"flag"
	"fmt"

	"github.com/deelawn/eth-beacon-time/beacon"
	slotLib "github.com/deelawn/eth-beacon-time/slot"
	timeLib "github.com/deelawn/eth-beacon-time/time"
)

var validChainIDs = map[string]struct{}{
	beacon.ChainIDMainnet: {},
	beacon.ChainIDPrater:  {},
}

func main() {

	chainID := flag.String("chain-id", beacon.ChainIDMainnet, "ID of the chain to use")
	slot := flag.Uint64("slot", 0, "the slot to get the time of")
	time := flag.Uint64("time", 0, "unix time to get the slot at")

	flag.Parse()

	// Formatting.
	fmt.Println("")

	if _, ok := validChainIDs[*chainID]; !ok {
		fmt.Printf("invalid chain-id %s\n", *chainID)
		return
	}

	if *slot == 0 && *time == 0 {
		fmt.Println("either slot or time must be specified")
		return
	}

	if *slot != 0 && *time != 0 {
		fmt.Println("cannot specify both slot AND time")
		return
	}

	if *slot != 0 {
		slotTime := slotLib.ToTime(beacon.GenesisTimes[*chainID], *slot)
		fmt.Println(slotTime)
		return
	}

	timeSlot := timeLib.ToSlot(beacon.GenesisTimes[*chainID], *time)
	fmt.Println(timeSlot)

}
