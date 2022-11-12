// Copyright 2022 Nukleros
// Copyright 2021-2022 VMware, Inc.
// SPDX-License-Identifier: MIT

package fuzz

import "github.com/nukleros/yaml-jsonpath/pkg/yamlpath"

// Fuzz allows go-fuzz to drive the lexer/parser.
func Fuzz(data []byte) int {
	p, err := yamlpath.NewPath(string(data))
	if err != nil {
		if p != nil {
			panic("Path != nil on error")
		}
		return 0
	}
	return 1
}
