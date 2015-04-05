package stitch_test

import "github.com/bmizerany/assert"
import "github.com/tj/stitch"
import "testing"

func TestFile(t *testing.T) {
	s, err := stitch.File("fixtures/a.sh")
	assert.Equal(t, nil, err)

	exp := `echo "from a"

# source fixtures/b.sh:
echo "from b"



# source fixtures/nested/c.sh:
echo "from c"

# source fixtures/nested/again/d.sh:
echo "from d"

# source fixtures/e.sh:
echo "from e"`

	assert.Equal(t, exp, s)
}
