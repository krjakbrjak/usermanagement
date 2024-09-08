package agent

import (
	"path/filepath"
	"testing"

	"github.com/krjakbrjak/usermanagement/generated"
	"github.com/stretchr/testify/assert"
)

func TestParsePAMConfig(t *testing.T) {
	pamConf := filepath.Join("testdata", "pam.conf")
	policy := &generated.PasswordPolicyResponse{}
	pamParseErr := ParsePAMConfig(pamConf, policy)
	if pamParseErr != nil {
		t.Fatalf("failed to read pam.conf.conf file: %v", pamParseErr)
	}
	assert.Equal(t, int32(8), policy.MinLength, "Wrong minlen!")
	assert.Equal(t, int32(0), policy.MaxDays, "Wrong minlen!")

	policy = &generated.PasswordPolicyResponse{}
	notExistParseErr := ParsePAMConfig("fake", policy)
	if notExistParseErr == nil {
		t.Fatalf("This test should fail")
	}
}
