package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"hash/crc32"
	"os"
	"strings"
)

var legacyPackage string

func main() {
	plugins := make(map[string]struct{})
	pgo := protogen.Options{
		ParamFunc: func(name, value string) error {
			switch name {
			case "plugin":
				for _, p := range strings.Split(value, "|") {
					plugins[p] = struct{}{}
				}
			}
			return nil
		},
	}

	legacyPackage, _ = os.LookupEnv("LEGACY_PACKAGE")
	if legacyPackage == "" {
		legacyPackage = "git.ronaksoft.com/river/server/msg/legacy"
	}

	cn := map[string]int64{}
	cs := map[int64]string{}
	pgo.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			g1 := plugin.NewGeneratedFile(fmt.Sprintf("%s.legacy.go", f.GeneratedFilenamePrefix), f.GoImportPath)
			GenConvertors(f, g1)
			for _, mt := range f.Messages {
				constructor := int64(crc32.ChecksumIEEE([]byte(mt.Desc.Name())))
				cn[string(mt.Desc.Name())] = constructor
				cs[constructor] = string(mt.Desc.Name())
			}
		}

		g2 := plugin.NewGeneratedFile("converter.go", "msg")
		GenMainConvertor(g2, cs)
		return nil
	})

	return
}

func GenMainConvertor(g *protogen.GeneratedFile, cs map[int64]string) {
	g.P("package ", "msg")
	g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "msg2",
		GoImportPath: protogen.GoImportPath(legacyPackage),
	})

	g.P("var cs = map[int64]func(d []byte) []byte{")
	for c, n := range cs {
		g.P(c, ": unmarshal", n, ",")
	}
	g.P("}")
	g.P("func ConvertToLegacy(c int64, d []byte) []date {")
	g.P("f := cs[c]")
	g.P("if f == nil {")
	g.P("panic(\"invalid constructor\")")
	g.P("}")
	g.P("return f(d)")
	g.P("}")
	g.P()
	for _, n := range cs {
		g.P("func unmarshal", n, "(d []byte) []byte {")
		g.P("x := &", n, "{}")
		g.P("_ = x.Unmarshal(d)")
		g.P("z := x.Convert()")
		g.P("b, _ := z.Marshal()")
		g.P("return b")
		g.P("}")
	}

}

// GenConvertors generates codes related for pooling of the messages
func GenConvertors(file *protogen.File, g *protogen.GeneratedFile) {
	g.P("package ", file.GoPackageName)
	g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "msg2",
		GoImportPath: protogen.GoImportPath(legacyPackage),
	})

	for _, mt := range file.Messages {
		mtName := mt.Desc.Name()
		g.P("func (x *", mtName, ") Convert () (z *legacy.", mtName, ") {")
		g.P("z = &legacy.", mtName, "{}")
		for _, ft := range mt.Fields {
			ftName := ft.Desc.Name()
			// ftPkg, _ := descName(file, g, ft.Desc.Message())
			switch ft.Desc.Cardinality() {
			case protoreflect.Repeated:
				switch ft.Desc.Kind() {
				case protoreflect.MessageKind:
					g.P("for _, item := range x.", ftName, "{")
					g.P("z.", ftName, "= append(z.", ftName, ", item.Convert())")
					g.P("}")
				default:
					g.P("z.", ftName, "= x.", ftName)
				}

			default:
				switch ft.Desc.Kind() {
				case protoreflect.MessageKind:
					g.P("z.", ftName, "= x.", ftName, ".Convert()")
				case protoreflect.EnumKind:
					g.P("z.", ftName, "= legacy.", ft.Desc.Enum().Name(), "(x.", ftName, ")")
				default:
					g.P("z.", ftName, "= x.", ftName)
				}
			}
		}
		g.P("return")
		g.P("}")
	}

	g.P()
}

// deskName returns the package and ident name
func descName(file *protogen.File, g *protogen.GeneratedFile, desc protoreflect.MessageDescriptor) (string, string) {
	if desc == nil {
		return "", ""
	}
	if string(desc.FullName().Parent()) == string(file.GoPackageName) {
		return "", string(desc.Name())
	} else {
		fd, ok := desc.ParentFile().Options().(*descriptorpb.FileOptions)
		if ok {
			g.QualifiedGoIdent(protogen.GoImportPath(fd.GetGoPackage()).Ident(string(desc.Name())))
		}
		return string(desc.ParentFile().Package()), string(desc.Name())
	}
}
