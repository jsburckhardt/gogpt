package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetVersion(t *testing.T) {
	myVersion := "edge"

	rootCmd.Version = myVersion
	setVersion()

	assert.Contains(t, rootCmd.VersionTemplate(), myVersion)
}
