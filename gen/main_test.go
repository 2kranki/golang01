// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Test main

package main

import (
	sharedData "github.com/2kranki/golang01/gen/shared"
	"testing"
)

func TestSetupDefns(t *testing.T) {
	var err error

	if err = SetupDefns("", "cmd"); err != nil {
		t.Errorf("SetupDefns() failed: %s\n", err)
	}
	if "cmd" != sharedData.Cmd() {
		t.Errorf("SetupDefns() cmd should be 'cmd', but is %s\n", sharedData.Cmd())
	}

	if err = SetupDefns("test/test01.exec.json.txt", "cmd"); err != nil {
		t.Errorf("SetupDefns() failed: %s\n", err)
	}
	if "cmd" != sharedData.Cmd() {
		t.Errorf("SetupDefns() cmd should be 'cmd', but is %s\n", sharedData.Cmd())
	}

}
