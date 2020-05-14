/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package disabled_test

import (
	"testing"
)

func TestDisabled(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Disabled Suite")
}
