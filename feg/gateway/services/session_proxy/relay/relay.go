/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package relay

import (
	"fmt"
	"strings"

	"magma/feg/gateway/registry"
	"magma/lte/cloud/go/protos"

	"google.golang.org/grpc"
)

type CloseableSessionProxyResponderClient struct {
	protos.SessionProxyResponderClient
	conn *grpc.ClientConn
}

func (client *CloseableSessionProxyResponderClient) Close() {
	client.conn.Close()
}

// Get a client to the local session manager client. To avoid leaking
// connections, defer Close() on the returned client.
func GetSessionProxyResponderClient(cloudRegistry registry.CloudRegistry) (*CloseableSessionProxyResponderClient, error) {
	conn, err := cloudRegistry.GetCloudConnection("feg_to_gw_relay")
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to gw relay: %s", err)
	}
	return &CloseableSessionProxyResponderClient{
		SessionProxyResponderClient: protos.NewSessionProxyResponderClient(conn),
		conn:                        conn,
	}, nil
}

func GetIMSIFromSessionID(sessionID string) (string, error) {
	split := strings.Split(sessionID, "-")
	if len(split) < 2 {
		return "", fmt.Errorf("Session ID %s does not match format 'IMSI-RandNum'", sessionID)
	}
	return split[0], nil
}
