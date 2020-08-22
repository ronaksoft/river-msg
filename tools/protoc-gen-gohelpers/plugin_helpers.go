package main

import (
	"fmt"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
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
	g *generator.Generator
}

func (g *GenPools) Name() string {
	return "gen-pools"
}

func (g *GenPools) Init(gen *generator.Generator) {
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
		constructor := crc32.ChecksumIEEE([]byte(*mt.Name))
		g.g.P(fmt.Sprintf("const C_%s int64 = %d", *mt.Name, constructor))
		initFunc.WriteString(fmt.Sprintf("ConstructorNames[%d] = \"%s\"\n", constructor, *mt.Name))
		g.g.P(fmt.Sprintf("type pool%s struct{", *mt.Name))
		g.g.In()
		g.g.P("pool sync.Pool")
		g.g.Out()
		g.g.P("}")
		g.g.P(fmt.Sprintf("func (p *pool%s) Get() *%s {", *mt.Name, *mt.Name))
		g.g.In()
		g.g.P(fmt.Sprintf("x, ok := p.pool.Get().(*%s)", *mt.Name))
		g.g.P("if !ok {")
		g.g.In()
		g.g.P(fmt.Sprintf("return &%s{}", *mt.Name))
		g.g.Out()
		g.g.P("}")
		g.g.P("return x")
		g.g.Out()
		g.g.P("}")
		g.g.P("", "")

		g.g.P(fmt.Sprintf("func (p *pool%s) Put(x *%s) {", *mt.Name, *mt.Name))
		g.g.In()
		for _, ft := range mt.Field {
			if *ft.Label == descriptor.FieldDescriptorProto_LABEL_OPTIONAL {
				g.g.P(zeroValue(file.GoPackageName(), ft.GetName(), ft.GetType(), ft.GetTypeName()))
			} else if *ft.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
				g.g.P(fmt.Sprintf("x.%s = x.%s[:0]", *ft.Name, *ft.Name))
			} else if *ft.Type == descriptor.FieldDescriptorProto_TYPE_BYTES {
				g.g.P(fmt.Sprintf("x.%s = x.%s[:0]", *ft.Name, *ft.Name))
			}
		}
		g.g.P("p.pool.Put(x)")
		g.g.Out()
		g.g.P("}")
		g.g.P("")
		g.g.P(fmt.Sprintf("var Pool%s = pool%s{}", *mt.Name, *mt.Name))
		g.g.P("")
		g.g.P(fmt.Sprintf("func Result%s(out *MessageEnvelope, res *%s) {", *mt.Name, *mt.Name))
		g.g.In()
		g.g.P(fmt.Sprintf("out.Constructor = C_%s", *mt.Name))
		g.g.P("protoSize := res.Size()")
		g.g.P("if protoSize > cap(out.Message) {")
		g.g.In()
		g.g.P("pbytes.Put(out.Message)")
		g.g.P("out.Message = pbytes.GetLen(protoSize)")
		g.g.Out()
		g.g.P("} else {")
		g.g.In()
		g.g.P("out.Message = out.Message[:protoSize]")
		g.g.Out()
		g.g.P("}")
		g.g.P("res.MarshalToSizedBuffer(out.Message)")
		g.g.Out()
		g.g.P("}")
	}
	initFunc.WriteString("}")
	g.g.P("")
	g.g.P(initFunc.String())

}

func zeroValue(gopkg string, n string,t descriptor.FieldDescriptorProto_Type, tn string) string {
	switch t {
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return fmt.Sprintf("x.%s = false", n)
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return fmt.Sprintf("x.%s = \"\"", n)
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		var ttn string
		tnp := strings.SplitN(tn[1:], ".", 2)
		if tnp[0] == gopkg {
			ttn = tnp[1]
		} else {
			ttn = tn[1:]
		}
		sb := strings.Builder{}
		sb.WriteString(fmt.Sprintf("if x.%s != nil {\n", n))
		sb.WriteString(fmt.Sprintf("*x.%s = %s{}", n, ttn))
		sb.WriteString("}\n")
		return sb.String()
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return fmt.Sprintf("x.%s = x.%s[:0]", n, n)
	default:
		return fmt.Sprintf("x.%s = 0", n)
	}
}

func (g *GenPools) GenerateImports(file *generator.FileDescriptor) {
	g.g.AddImport("sync")
	g.g.AddImport("git.ronaksoft.com/river")
}
