package main

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"hash/crc32"
	"strings"
)

/*
   Creation Time: 2019 - Nov - 29
   Created by:  (ehsan)
   Maintainers:
      1.  Ehsan N. Moosa (E2)
   Auditor: Ehsan N. Moosa (E2)
   Copyright Ronak Software Group 2018
*/

type GenPools struct {
	g  *generator.Generator
	cn map[string]int64
	cs map[int64]string
}

func (g *GenPools) Name() string {
	return "gen-pools"
}

func (g *GenPools) Init(gen *generator.Generator) {
	g.cn = make(map[string]int64)
	g.cs = make(map[int64]string)
	if gen == nil {
		g.g = generator.New()
	} else {
		g.g = gen
	}
}

func (g *GenPools) Generate(file *generator.FileDescriptor) {
	g.g.AddImport("sync")
	g.g.AddImport("github.com/gobwas/pool/pbytes")
	initFunc := strings.Builder{}
	initFunc.WriteString("func init() {\n")
	for _, mt := range file.MessageType {
		constructor := int64(crc32.ChecksumIEEE([]byte(*mt.Name)))
		g.cn[*mt.Name] = constructor
		g.cs[constructor] = *mt.Name
	}
}

// func zeroValue(t *descriptor.FieldDescriptorProto_Type) string {
// 	switch *t {
// 	case descriptor.FieldDescriptorProto_TYPE_BOOL:
// 		return "false"
// 	case descriptor.FieldDescriptorProto_TYPE_STRING:
// 		return "\"\""
// 	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
// 		return "nil"
// 	case descriptor.FieldDescriptorProto_TYPE_BYTES:
// 		return "nil"
// 	default:
// 		return "0"
// 	}
// }

func (g *GenPools) GenerateImports(file *generator.FileDescriptor) {
	//
	// g.g.AddImport("sync")
	// g.g.AddImport("git.ronaksoftware.com/river")
}
