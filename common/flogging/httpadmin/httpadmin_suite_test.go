/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package httpadmin_test

import (
	"testing"
)

func TestHttpadmin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Httpadmin Suite")
}
