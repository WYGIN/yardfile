package pkg

import (
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

func init()  {
	_ = llb.State{}
	instructions.Parse(nil, nil)
	parser.Parse(nil)
}