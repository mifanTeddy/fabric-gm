/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package privdata

import (
	"testing"

	"github.com/hyperledger/fabric-gm/common/cauthdsl"
	"github.com/hyperledger/fabric-gm/msp"
	"github.com/hyperledger/fabric-gm/protos/common"
	"github.com/stretchr/testify/assert"
)

func TestMembershipInfoProvider(t *testing.T) {
	peerSelfSignedData := common.SignedData{
		Identity:  []byte("peer0"),
		Signature: []byte{1, 2, 3},
		Data:      []byte{4, 5, 6},
	}

	identityDeserializer := func(chainID string) msp.IdentityDeserializer {
		return &mockDeserializer{}
	}

	// verify membership provider returns true
	membershipProvider := NewMembershipInfoProvider(peerSelfSignedData, identityDeserializer)
	res, err := membershipProvider.AmMemberOf("test1", getAccessPolicy([]string{"peer0", "peer1"}))
	assert.True(t, res)
	assert.Nil(t, err)

	// verify membership provider returns false
	res, err = membershipProvider.AmMemberOf("test1", getAccessPolicy([]string{"peer2", "peer3"}))
	assert.False(t, res)
	assert.Nil(t, err)

	// verify membership provider returns nil and error when collection policy config is nil
	res, err = membershipProvider.AmMemberOf("test1", nil)
	assert.False(t, res)
	assert.Error(t, err)
	assert.Equal(t, "Collection policy config is nil", err.Error())
}

func getAccessPolicy(signers []string) *common.CollectionPolicyConfig {
	var data [][]byte
	for _, signer := range signers {
		data = append(data, []byte(signer))
	}
	policyEnvelope := cauthdsl.Envelope(cauthdsl.Or(cauthdsl.SignedBy(0), cauthdsl.SignedBy(1)), data)
	return createCollectionPolicyConfig(policyEnvelope)
}
