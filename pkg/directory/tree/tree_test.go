package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivanilves/travelgrunt/pkg/config"
	"github.com/ivanilves/travelgrunt/pkg/directory"
)

const (
	fixturePath = "../../../fixtures/directory"
)

var (
	mock = mockTree()
)

func mockTree() Tree {
	cfg := config.DefaultConfig()

	_, paths, _ := directory.Collect(fixturePath, cfg)

	return NewTree(paths)
}

func TestSortedKeys(t *testing.T) {
	assert := assert.New(t)

	input := map[string]string{
		"prod": "terragrunt/prod",
		"dev":  "terragrunt/dev",
	}

	expected := []string{"dev", "prod"}

	assert.Equal(expected, sortedKeys(input))
}

func TestNewTree(t *testing.T) {
	assert := assert.New(t)

	assert.NotNil(mock)
}

func TestLevelCount(t *testing.T) {
	assert := assert.New(t)

	expected := 5

	assert.Equal(expected, mock.LevelCount())
}

func TestLevelItems(t *testing.T) {
	assert := assert.New(t)

	assert.Nil(mock.levelItems(10))

	items := mock.levelItems(0)

	assert.NotNil(items)
	assert.Equal(1, len(items))
	assert.Equal("terragrunt", items["terragrunt"])
}

func TestLevelChildItems(t *testing.T) {
	assert := assert.New(t)

	assert.Nil(mock.LevelChildItems(10, "whatever"))
	assert.Equal(map[string]string{}, mock.LevelChildItems(4, "terragrunt/prod/region-1/k8s/foo"))
	assert.Equal(mock.levelItems(0), mock.LevelChildItems(-1, "ignored-if-minus-one-passed"))

	items := mock.LevelChildItems(0, "terragrunt")
	expected := map[string]string{"dev": "terragrunt/dev", "prod": "terragrunt/prod"}

	assert.NotNil(items)
	assert.Equal(2, len(items))
	assert.Equal(expected, items)
}

func TestLevelChildItemsTerminalAndWithChildren(t *testing.T) {
	assert := assert.New(t)

	items := mock.LevelChildItems(3, "terragrunt/dev/region-1/rds")
	expected := map[string]string{
		".":   "terragrunt/dev/region-1/rds",
		"bar": "terragrunt/dev/region-1/rds/bar",
		"baz": "terragrunt/dev/region-1/rds/baz",
		"foo": "terragrunt/dev/region-1/rds/foo",
	}

	assert.NotNil(items)
	assert.Equal(4, len(items))
	assert.Equal(expected, items)
}

func TestLevelChildNames(t *testing.T) {
	assert := assert.New(t)

	items := mock.LevelChildNames(0, "terragrunt")
	expected := []string{"dev", "prod"}

	assert.NotNil(items)
	assert.Equal(2, len(items))
	assert.Equal(expected, items)
}
