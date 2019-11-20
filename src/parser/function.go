// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseFunction() (ast.Node, error) {
	node := ast.Function{}.Init(p.peek.Line, p.peek.Type)
	p.next() // 'fn'

	fnName, err := p.parseIdentifier()
	if err != nil {
		return nil, err
	}
	node.SetName(fnName.(ast.Identifier))

	err = p.require(p.peek.Type, types.LPAREN)
	if err != nil {
		return nil, err
	}
	p.next()

	var params []ast.Identifier = make([]ast.Identifier, 0)
	for p.peek.Type == types.IDENT {
		ident, err := p.parseIdentifier()
		if err != nil {
			return nil, err
		}
		params = append(params, ident.(ast.Identifier))
	}
	node.SetParams(params)

	err = p.require(p.peek.Type, types.RPAREN)
	if err != nil {
		return nil, err
	}
	p.next()

	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}
	node.SetBody(body)

	err = p.require(p.peek.Type, types.EOL)
	if err != nil {
		return nil, err
	}
	p.next()

	return node, nil
}
