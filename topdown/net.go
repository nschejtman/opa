// Copyright 2021 The OPA Authors.  All rights reserved.
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package topdown

import (
	"fmt"
	"github.com/open-policy-agent/opa/ast"
)

type lookupIPAddrCacheKey string

func builtinLookupIPAddr(bctx BuiltinContext, operands []*ast.Term, iter func(*ast.Term) error) error {
	return fmt.Errorf("builtinLookupIPAddr not implemented")
}

func init() {
	RegisterBuiltinFunc(ast.NetLookupIPAddr.Name, builtinLookupIPAddr)
}
