/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package etcdraft_test

import (
	"testing"
)

var testingInstance *testing.T

func TestEtcdraft(t *testing.T) {
	testingInstance = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Etcdraft Suite")
}
