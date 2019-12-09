package main

import (
	"bytes"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
	"hash/crc32"
	"text/template"
)

func main() {
	req := command.Read()
	files := req.GetProtoFile()
	files = vanity.FilterFiles(files, vanity.NotGoogleProtobufDescriptorProto)

	cn := map[string]int64{}
	cs := map[int64]string{}
	for _, file := range files {
		for _, mt := range file.MessageType {
			constructor := int64(crc32.ChecksumIEEE([]byte(*mt.Name)))
			cn[*mt.Name] = constructor
			cs[constructor] = *mt.Name
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

	resp := &plugin.CodeGeneratorResponse{}
	fileName := "constructors.json"
	outStr := out.String()
	resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
		Name:           &fileName,
		Content:        &outStr,
		InsertionPoint: nil,
	})
	command.Write(resp)

	// vanity.ForEachFile(files, vanity.TurnOnMarshalerAll)
	// vanity.ForEachFile(files, vanity.TurnOnSizerAll)
	// vanity.ForEachFile(files, vanity.TurnOnUnmarshalerAll)
	// resp := command.GeneratePlugin(req, &GenPools{}, "_con.go")
	// command.Read()

	return
}
