//go:build tools
// +build tools

// See https://github.com/golang/go/issues/26366
package generator

import (
	_ "github.com/polarlightsllc/prisma-client-go-1/generator/templates"
	_ "github.com/polarlightsllc/prisma-client-go-1/generator/templates/actions"
)
