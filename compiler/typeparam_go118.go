//go:build go1.18
// +build go1.18

package compiler

import (
	"go/ast"
	"go/types"
)

func hasTypeParam(t types.Type) bool {
	switch t := t.(type) {
	case *types.TypeParam:
		return true
	case *types.Named:
		return t.TypeParams() != nil
	case *types.Signature:
		return t.TypeParams() != nil
	}
	return false
}

func funcHasTypeParam(t *ast.FuncType) bool {
	return t.TypeParams != nil
}

func namedHasTypeParam(t *types.Named) bool {
	return t.TypeParams() != nil
}
