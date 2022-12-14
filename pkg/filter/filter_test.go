package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	assert := assert.New(t)

	assert.Nil(Validate([]string{}))
	assert.Nil(Validate([]string{"eu", "prod"}))
	assert.NotNil(Validate([]string{"us", "z"}))
}

func TestApply(t *testing.T) {
	assert := assert.New(t)

	ss := []string{
		"prod/us-east-1/rds",
		"prod/us-east-2/rds",
		"prod/us-east-1/ec2",
		"dev/us-east-1/rds",
		"dev/eu-west-1/rds",
		"dev/us-east-1/ec2",
	}

	matches := []string{"rds", "east-1"}

	expected := []string{"dev/us-east-1/rds", "prod/us-east-1/rds"}

	assert.Equal(expected, Apply(ss, matches))
}

func TestApplyEmptyFilter(t *testing.T) {
	assert := assert.New(t)

	ss := []string{
		"prod/us-east-1/rds",
		"prod/us-east-2/rds",
		"prod/us-east-1/ec2",
		"dev/us-east-1/rds",
		"dev/eu-west-1/rds",
		"dev/us-east-1/ec2",
	}

	matches := []string{}

	expected := []string{
		"dev/eu-west-1/rds",
		"dev/us-east-1/ec2",
		"dev/us-east-1/rds",
		"prod/us-east-1/ec2",
		"prod/us-east-1/rds",
		"prod/us-east-2/rds",
	}

	assert.Equal(expected, Apply(ss, matches))
}
