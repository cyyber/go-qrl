package types

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

// TestGeneratedHeaderRLPIsFresh guards against hand edits to gen_header_rlp.go
// by checking that the committed file matches what rlpgen produces from the
// current Header definition.
func TestGeneratedHeaderRLPIsFresh(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping code generation check in short mode")
	}
	want, err := os.ReadFile("gen_header_rlp.go")
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "run", "../../rlp/rlpgen", "-type", "Header")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	got, err := cmd.Output()
	if err != nil {
		t.Fatalf("rlpgen failed: %v\n%s", err, stderr.String())
	}
	if !bytes.Equal(got, want) {
		t.Fatal("gen_header_rlp.go is out of date; run `go generate ./core/types`")
	}
}
