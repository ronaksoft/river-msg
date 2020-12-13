package main

import (
	"bytes"
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
				constructor := int64(crc32.ChecksumIEEE([]byte(mt.Desc.FullName())))
				cn[string(mt.Desc.Name())] = constructor
				cs[constructor] = string(mt.Desc.FullName())
			}
		}

		t := template.Must(template.New("t1").Parse(`
	{
	    "ConstructorsByName": {
	    {{range $k,$v := .}}    "{{$k}}": {{$v}},
		{{end -}}
		},
		"ConstructorsByValue": {
		{{range $k,$v := .}}    "{{$v}}": "{{$k}}",
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

// func main() {
// 	req := command.Read()
// 	files := req.GetProtoFile()
// 	files = vanity.FilterFiles(files, vanity.NotGoogleProtobufDescriptorProto)
//
// 	cn := map[string]int64{}
// 	cs := map[int64]string{}
// 	for _, file := range files {
// 		for _, mt := range file.MessageType {
// 			constructor := int64(crc32.ChecksumIEEE([]byte(*mt.Name)))
// 			cn[*mt.Name] = constructor
// 			cs[constructor] = *mt.Name
// 		}
// 	}
//
// 	t := template.Must(template.New("t1").Parse(`
// {
//     "ConstructorsByName": {
//     {{range $k,$v := .}}    "{{$k}}": {{$v}},
// 	{{end -}}
// 	},
// 	"ConstructorsByValue": {
// 	{{range $k,$v := .}}    {{$v}}: "{{$k}}",
// 	{{end -}}
// 	}
// }
// `))
//
// 	out := &bytes.Buffer{}
// 	err := t.Execute(out, cn)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	resp := &plugin.CodeGeneratorResponse{}
// 	fileName := "constructors.json"
// 	outStr := out.String()
// 	resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
// 		Name:           &fileName,
// 		Content:        &outStr,
// 		InsertionPoint: nil,
// 	})
// 	command.Write(resp)
//
// 	// vanity.ForEachFile(files, vanity.TurnOnMarshalerAll)
// 	// vanity.ForEachFile(files, vanity.TurnOnSizerAll)
// 	// vanity.ForEachFile(files, vanity.TurnOnUnmarshalerAll)
// 	// resp := command.GeneratePlugin(req, &GenPools{}, "_con.go")
// 	// command.Read()
//
// 	return
// }
