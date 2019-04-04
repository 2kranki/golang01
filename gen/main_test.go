// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Test main

package main

import "testing"

func TestSetupDefns(t *testing.T) {
	var err error

	err = SetupDefns("", "cmd")
	if err != nil {
		t.Errorf("SetupDefns() failed: %s\n", err)
	}

	err = SetupDefns("test.exec.json.txt", "cmd")
	if err != nil {
		t.Errorf("SetupDefns() failed: %s\n", err)
	}

}
