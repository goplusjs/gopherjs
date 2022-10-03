//go:build !go1.18
// +build !go1.18

package typeparams

import (
	"go/ast"
	"go/types"
)

func HasTypeParam(t types.Type) bool {
	return false
}

func NamedHasTypeParam(t *types.Named) bool {
	return false
}

func FuncTypeHasTypeParam(t *ast.FuncType) bool {
	return false
}

func TypeSpecRemoveTypeParam(spec *ast.TypeSpec) {
}
