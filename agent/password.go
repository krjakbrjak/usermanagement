package agent

//go:generate go run ../tools/generate.go

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/krjakbrjak/usermanagement/generated"
)

type Agent struct {
	generated.UnimplementedPasswordPolicyServiceServer
}

func (s *Agent) GetPasswordPolicy(ctx context.Context, req *empty.Empty) (*generated.PasswordPolicyResponse, error) {
	policy := &generated.PasswordPolicyResponse{}
	installedErr := CheckPAMPwquality()
	if installedErr != nil {
		return nil, installedErr
	}
	parseErr := ParsePAMConfig("/etc/pam.d/common-password", policy)
	if parseErr != nil {
		return nil, parseErr
	}
	return policy, nil
}
