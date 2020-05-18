package opa

import (
	"context"
	"github.com/open-policy-agent/opa/rego"
	"github.com/sirupsen/logrus"
)

type RegoInput struct {
	Path   string   `json:"path"`
	Method string   `json:"method"`
	Token  string   `json:"token"`
}

func GetDecision (requestRegoInput RegoInput) bool {
	var err error
	ctx := context.Background()

	authorizationRego := rego.New(
		rego.Query("data.authorization.isAccessGranted"),
		rego.Input(requestRegoInput),
		rego.Load([]string{"../../pkg/rbac/opa/authorization.rego",
			"../../pkg/rbac/api/middleware/authorizationCache.json"}, nil))

	regoResult, err := authorizationRego.Eval(ctx)
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	return regoResult[0].Expressions[0].Value.(bool)
}