package jsonpath

import (
	"errors"

	"github.com/pkg-base/libopenapi/internal/jsonpath/config"
	"github.com/pkg-base/libopenapi/internal/jsonpath/token"
	"github.com/pkg-base/yaml"
)

func NewPath(input string, opts ...config.Option) (*JSONPath, error) {
	tokenizer := token.NewTokenizer(input, opts...)
	tokens := tokenizer.Tokenize()
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Token == token.ILLEGAL {
			return nil, errors.New(tokenizer.ErrorString(&tokens[i], "unexpected token"))
		}
	}
	parser := newParserPrivate(tokenizer, tokens, opts...)
	err := parser.parse()
	if err != nil {
		return nil, err
	}
	return parser, nil
}

func (p *JSONPath) Query(root *yaml.Node) []*yaml.Node {
	return p.ast.Query(root, root)
}

func (p *JSONPath) String() string {
	if p == nil {
		return ""
	}
	return p.ast.ToString()
}
