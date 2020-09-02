package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twpayne/go-vfs/vfst"
)

func TestDiffCmd(t *testing.T) {
	fs, cleanup, err := vfst.NewTestFS(map[string]interface{}{
		"/home/user": map[string]interface{}{
			".hushlogin": "",
			".local/share/chezmoi/empty_dot_hushlogin": "",
		},
	})
	require.NoError(t, err)
	defer cleanup()

	stdout := &bytes.Buffer{}
	c := newTestConfig(t, fs, withStdout(stdout))
	assert.NoError(t, c.runDiffCmd(nil, nil))
	assert.Empty(t, stdout.String())
}
