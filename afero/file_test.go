package afero

import (
	"github.com/spf13/afero"
	"testing"
)

func Test_file(t *testing.T) {
	afero.NewMemMapFs()
}
