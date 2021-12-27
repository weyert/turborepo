package fs

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	vectors := []struct {
		file string // Test input file
		want YarnLockfile
	}{
		{
			file: "testdata/yarn_normal.lock",
			want: yarnNormal,
		},

		// // {
		// // 	file: "testdata/yarn_v2_normal.lock",
		// // 	want: yarnNormalV2,
		// // },

		// {
		// 	file: "testdata/yarn_normal-windows.lock",
		// 	want: yarnNormalWindows,
		// },
	}

	for _, v := range vectors {
		t.Run(path.Base(v.file), func(t *testing.T) {
			lockfile, err := parseYarnLockfile(v.file)
			require.NoError(t, err)

			assert.Equal(t, &v.want, lockfile)
		})
	}
}
