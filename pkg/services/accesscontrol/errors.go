package accesscontrol

import (
	"errors"
	"fmt"
)

var (
	ErrFixedRolePrefixMissing = errors.New("fixed role should be prefixed with '" + FixedRolePrefix + "'")
	ErrInvalidBuiltinRole     = errors.New("built-in role is not valid")
	ErrInvalidScope           = errors.New("invalid scope")
	ErrResolverNotFound       = errors.New("no resolver found")
)

type ErrorInvalidRole struct{}

func (e *ErrorInvalidRole) Error() string {
	return "role is invalid"
}

type ErrorRolePrefixMissing struct {
	Role     string
	Prefixes []string
}

func (e *ErrorRolePrefixMissing) Error() string {
	return fmt.Sprintf("expected role '%s' to be prefixed with any of '%v'", e.Role, e.Prefixes)
}

func (e *ErrorRolePrefixMissing) Unwrap() error {
	return &ErrorInvalidRole{}
}

type ErrorActionPrefixMissing struct {
	Action   string
	Prefixes []string
}

func (e *ErrorActionPrefixMissing) Error() string {
	return fmt.Sprintf("expected action '%s' to be prefixed with any of '%v'", e.Action, e.Prefixes)
}

func (e *ErrorActionPrefixMissing) Unwrap() error {
	return &ErrorInvalidRole{}
}

type ErrorInvalidEvaluationRequest struct {
	Action    string
	Resource  string
	Attribute string
	UIDs      []string
}

func (e *ErrorInvalidEvaluationRequest) Error() string {
	switch {
	case e.Action == "" && e.Resource == "" && e.Attribute == "" && len(e.UIDs) == 0:
		return "no filtering field is set"
	case e.Action == "":
		return "all resource filtering fields must be set"
	default:
		return "none or all resource filtering fields must be set"
	}
}
