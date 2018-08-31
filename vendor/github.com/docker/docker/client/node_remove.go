package client

import (
	"context"
	"net/url"

	"github.com/docker/docker/api/types"
)

// NodeRemove removes a Node.
func (cli *Client) NodeRemove(ctx context.Context, nodeID string, options types.NodeRemoveOptions) error {
	query := url.Values{}
	if options.Force {
		query.Set("force", "1")
	}

	resp, err := cli.delete(ctx, "/nodes/"+nodeID, query, nil)
	ensureReaderClosed(resp)
	return wrapResponseError(err, resp, "node", nodeID)
}
