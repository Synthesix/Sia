package client

import (
	"fmt"

	"github.com/Synthesix/Sia/node/api"
	"github.com/Synthesix/Sia/types"
)

// ConsensusGet requests the /consensus api resource
func (c *Client) ConsensusGet() (cg api.ConsensusGET, err error) {
	err = c.get("/consensus", &cg)
	return
}

// ConsensusBlocksIDGet requests the /consensus/blocks api resource
func (c *Client) ConsensusBlocksIDGet(id types.BlockID) (block types.Block, err error) {
	err = c.get("/consensus/blocks?id="+id.String(), &block)
	return
}

// ConsensusBlocksHeightGet requests the /consensus/blocks api resource
func (c *Client) ConsensusBlocksHeightGet(height types.BlockHeight) (block types.Block, err error) {
	err = c.get("/consensus/blocks?height="+fmt.Sprint(height), &block)
	return
}
