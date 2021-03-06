package node

import (
	"fmt"

	"github.com/elastos/Elastos.ELA/protocol"

	"github.com/elastos/Elastos.ELA.Utility/p2p"
)

// Extra net nodes will connect through the specific NodeOpenPort
// only limited messages can go through, such as filterload/getblocks/getdata etc.
// Otherwise error will be returned
func FilterMessage(node protocol.Noder, msgType string) error {
	if node.IsFromExtraNet() {
		// Only cased message types can go through
		switch msgType {
		case p2p.CmdVersion:
		case p2p.CmdVerAck:
		case p2p.CmdGetAddr:
		case p2p.CmdPing:
		case p2p.CmdPong:
		case p2p.CmdFilterLoad:
		case p2p.CmdGetBlocks:
		case p2p.CmdGetData:
		case p2p.CmdTx:
		case p2p.CmdMemPool:
		default:
			return fmt.Errorf("unsupport messsage for extra net node %s", msgType)
		}
	}
	return nil
}
