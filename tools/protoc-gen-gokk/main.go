package main

import (
	"bytes"
	"fmt"
	_ "github.com/ronaksoft/rony"
	"google.golang.org/protobuf/compiler/protogen"
	"hash/crc32"
	"html/template"
	"strings"
)

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

	cn := map[string]int64{}
	cs := map[int64]string{}
	pgo.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			// reset the global model and fill with the new data
			for _, mt := range f.Messages {
				constructor := int64(crc32.ChecksumIEEE([]byte(mt.Desc.Name())))
				cn[string(mt.Desc.Name())] = constructor
				cs[constructor] = string(mt.Desc.Name())
			}
			for _, s := range f.Services {
				for _, m := range s.Methods {
					methodName := fmt.Sprintf("%s%s", s.Desc.Name(), m.Desc.Name())
					constructor := int64(crc32.ChecksumIEEE([]byte(methodName)))
					cn[methodName] = constructor
					cs[constructor] = methodName
				}
			}
		}

		t := template.Must(template.New("t1").Parse(`
	{
	    "ConstructorsByName": {
	    {{range $k,$v := .}}    "{{$k}}": {{$v}},
		{{end -}}
		},
		"ConstructorsByValue": {
		{{range $k,$v := .}}    {{$v}}: "{{$k}}",
		{{end -}}
		}
	}
	`))

		out := &bytes.Buffer{}
		err := t.Execute(out, cn)
		if err != nil {
			panic(err)
		}

		gf := plugin.NewGeneratedFile("constructors.json", "msg")
		_, err = gf.Write(out.Bytes())
		return err
	})

	return
}
