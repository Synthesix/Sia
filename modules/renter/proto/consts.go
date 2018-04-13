package proto

import (
	"time"

	"github.com/Synthesix/Sia/build"
	"github.com/Synthesix/Sia/crypto"
	"github.com/Synthesix/Sia/modules"
)

const (
	// contractExtension is the extension given to contract files.
	contractExtension = ".contract"
)

var (
	// connTimeout determines the number of seconds the dialer will wait
	// for a connect to complete
	connTimeout = build.Select(build.Var{
		Dev:      10 * time.Second,
		Standard: 60 * time.Second,
		Testing:  5 * time.Second,
	}).(time.Duration)

	// hostPriceLeeway is the amount of flexibility we give to hosts when
	// choosing how much to pay for file uploads. If the host does not have the
	// most recent block yet, the host will be expecting a slightly larger
	// payment.
	//
	// TODO: Due to the network connectivity issues that v1.3.0 introduced, we
	// had to increase the amount moderately because hosts would not always be
	// properly connected to the peer network, and so could fall behind on
	// blocks. Once enough of the network has upgraded, we can move the number
	// to '0.003' for 'Standard'.
	hostPriceLeeway = build.Select(build.Var{
		Dev:      0.05,
		Standard: 0.01,
		Testing:  0.002,
	}).(float64)

	// sectorHeight is the height of a Merkle tree that covers a single
	// sector. It is log2(modules.SectorSize / crypto.SegmentSize)
	sectorHeight = func() uint64 {
		height := uint64(0)
		for 1<<height < (modules.SectorSize / crypto.SegmentSize) {
			height++
		}
		return height
	}()
)
