package nodejs

import (
	"testing"
)

func TestNodeSdk_Versions(t *testing.T) {
	sdk := &NodeSdk{}
	sdk.Versions()
}

func TestNodeSdk_Install(t *testing.T) {
	sdk := &NodeSdk{}
	sdk.Install("v9.9.0")
}

func TestNodeSdk_LocalVersions(t *testing.T) {
	sdk := &NodeSdk{}
	sdk.LocalVersions()
}
func TestNodeSdk_UseVersion(t *testing.T) {
	sdk := &NodeSdk{}
	sdk.UseVersion("v9.9.0")
}
