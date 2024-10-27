package rule

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fixturePath = "../../../fixtures/config/rule"
)

func TestPathPrefixed(t *testing.T) {
	assert := assert.New(t)

	assert.False(Rule{Prefix: "./prefix"}.pathPrefixed("/etc/passwd"))
	assert.True(Rule{Prefix: ""}.pathPrefixed("/etc/passwd"))
	assert.True(Rule{Prefix: "./prefix"}.pathPrefixed("./prefix/file"))
}

func TestNormalizeNameEx(t *testing.T) {
	assert := assert.New(t)

	r := Rule{NameEx: "*.inc.php"}

	assert.Equal(r.normalizeNameEx(), "^.*\\.inc\\.php$")
}

func TestNameMatched(t *testing.T) {
	assert := assert.New(t)

	files, err := os.ReadDir(fixturePath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		assert.True(Rule{}.nameMatched(file))
		assert.True(Rule{NameEx: ".*"}.nameMatched(file))
		assert.False(Rule{NameEx: "some-nonsense"}.nameMatched(file))
	}
}
