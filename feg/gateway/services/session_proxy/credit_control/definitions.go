/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package credit_control

type CreditRequestType uint8

const (
	CRTInit      CreditRequestType = 0x01
	CRTUpdate    CreditRequestType = 0x02
	CRTTerminate CreditRequestType = 0x03

	SessionProxyServiceName = "session_proxy"
)

type RequestKeyNamespace int

const (
	None RequestKeyNamespace = iota
	Gx
	Gy
)

type SubscriptionIDType uint8

type GrantedServiceUnit struct {
	TotalOctets  *uint64 `avp:"CC-Total-Octets"`
	InputOctets  *uint64 `avp:"CC-Input-Octets"`
	OutputOctets *uint64 `avp:"CC-Output-Octets"`
}

const (
	EndUserE164 SubscriptionIDType = 0x0
	EndUserIMSI SubscriptionIDType = 0x1
)

type RequestKey struct {
	Namespace     RequestKeyNamespace
	SessionID     string
	RequestNumber uint32
}

// GetRequestKey generates request tracking key based on Namespace, session ID & request number
func GetRequestKey(ns RequestKeyNamespace, sessionID string, requestNumber uint32) RequestKey {
	return RequestKey{
		Namespace:     ns,
		SessionID:     sessionID,
		RequestNumber: requestNumber,
	}
}
