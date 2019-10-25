package nodejs

import "testing"

func TestNodeSdk_Versions(t *testing.T) {
	sdk := &NodeSdk{}
	sdk.Versions()
}

func TestNodeSdk_Install(t *testing.T) {
	sdk := &NodeSdk{}
	sdk.Install("v9.9.0")
}
