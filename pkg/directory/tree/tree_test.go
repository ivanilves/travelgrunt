package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivanilves/ttg/pkg/directory"
)

const (
	fixturePath = "../../../fixtures/directory"
)

var (
	mock = mockTree()
)

func mockTree() Tree {
	_, paths, _ := directory.Collect(fixturePath)

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

	assert.Nil(mock.LevelItems(10))

	items := mock.LevelItems(0)

	assert.NotNil(items)
	assert.Equal(1, len(items))
	assert.Equal("terragrunt", items["terragrunt"])
}

func TestChildItems(t *testing.T) {
	assert := assert.New(t)

	assert.Nil(mock.ChildItems(10, "whatever"))
	assert.Equal(map[string]string{}, mock.ChildItems(4, "terragrunt/prod/region-1/k8s/foo"))
	assert.Equal(mock.LevelItems(0), mock.ChildItems(-1, "ignored-if-minus-one-passed"))

	items := mock.ChildItems(0, "terragrunt")
	expected := map[string]string{"dev": "terragrunt/dev", "prod": "terragrunt/prod"}

	assert.NotNil(items)
	assert.Equal(2, len(items))
	assert.Equal(expected, items)
}

func TestChildNames(t *testing.T) {
	assert := assert.New(t)

	items := mock.ChildNames(0, "terragrunt")
	expected := []string{"dev", "prod"}

	assert.NotNil(items)
	assert.Equal(2, len(items))
	assert.Equal(expected, items)
}

func TestHasChildren(t *testing.T) {
	assert := assert.New(t)

	assert.True(mock.HasChildren(-1, ""))
	assert.True(mock.HasChildren(-1, "whatever"))

	assert.True(mock.nodeExists("terragrunt"))
	assert.True(mock.HasChildren(0, "terragrunt"))

	assert.True(mock.nodeExists("terragrunt/prod/region-1/k8s"))
	assert.True(mock.HasChildren(3, "terragrunt/prod/region-1/k8s"))

	assert.False(mock.nodeExists("i-do-not-exist"))
	assert.False(mock.HasChildren(0, "i-do-not-exist"))

	assert.True(mock.nodeExists("terragrunt/prod/region-1/k8s/foo"))
	assert.False(mock.HasChildren(4, "terragrunt/prod/region-1/k8s/foo"))
}
