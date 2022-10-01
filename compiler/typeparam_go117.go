//go:build !go1.18
// +build !go1.18

package compiler

import (
	"go/ast"
	"go/types"
)

func hasTypeParam(t types.Type) bool {
	return false
}

func funcHasTypeParam(t *ast.FuncType) bool {
	return false
}

func namedHasTypeParam(t *types.Named) bool {
	return false
}
