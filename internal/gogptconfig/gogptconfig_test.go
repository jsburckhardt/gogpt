package gogptconfig_test

import (
	"gogpt/internal/gogptconfig"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVars(t *testing.T) {

	_ = os.Setenv("TESTV1", "true")
	_ = os.Setenv("TESTV2", "true")
	_ = os.Setenv("TESTV2.1", "false")
	_ = os.Setenv("TESTV3", "TESTV3")
	_ = os.Setenv("TESTV4", "TESTV4")

	conf := gogptconfig.GetConfig()

	testv1 := conf.GetBool("TESTV1")
	assert.True(t, testv1)

	testv2 := conf.GetBool("testv2")
	assert.True(t, testv2)

	testv22 := conf.GetBool("TESTV2.1")
	assert.False(t, testv22)

	testv3 := conf.GetString("TESTV3")
	assert.Equal(t, "TESTV3", testv3)

	testv4 := conf.GetString("TESTV4")
	assert.Equal(t, "TESTV4", testv4)

	_ = os.Setenv("TEST", "TEST Override")

	testv6 := conf.GetString("TEST")
	assert.Equal(t, "TEST Override", testv6)
}
