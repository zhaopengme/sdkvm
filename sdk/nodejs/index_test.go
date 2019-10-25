package nodejs

import "testing"

func TestNodeSdk_Versions(t *testing.T) {
	sdk := &NodeSdk{}
	sdk.Versions()
}
