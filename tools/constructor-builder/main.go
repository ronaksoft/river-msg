package main

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	messages map[uint32]string
)

var (
	constructorsTemplate = `
const (
    {{- range $constructor, $messageName := .}}
    C_{{- $messageName }} int64 = {{ $constructor -}}
    {{- end }}
)

var ConstructorNames = map[int64]string{
    {{- range $constructor, $messageName := .}}
    {{ $constructor }}: "{{$messageName}}",
    {{- end }}
}
`
	poolTemplate = `
import (
    "sync"
)
var (
{{range $constructor, $messageName := .}}
Pool{{$messageName}} = sync.Pool{
        New: func() interface{} {
            m := new({{$messageName}})
            return m
        },
    }
{{- end}}
)
`

	resultTemplate = `
import (
	"github.com/gobwas/pool/pbytes"
)
{{range $constructor, $messageName := .}}
func Result{{$messageName}}(out *MessageEnvelope, res *{{$messageName}}) {
	out.Constructor = C_{{$messageName}}
	pbytes.Put(out.Message)
	out.Message = pbytes.GetLen(res.Size())
	res.MarshalTo(out.Message)
}
{{- end}}

`
)

func main() {
	var path string
	var packageName string
	if len(os.Args) > 1 {
		path = strings.Trim(os.Args[1], "/ ")
	} else {
		path = "."
	}
	packageName = os.Args[2]
	constructorsTemplate = fmt.Sprintf("package %s \r\n %s", packageName, constructorsTemplate)
	poolTemplate = fmt.Sprintf("package %s \r\n %s", packageName, poolTemplate)
	resultTemplate = fmt.Sprintf("package %s \r\n %s", packageName, resultTemplate)

	messages = make(map[uint32]string)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".proto") {
			ScanFile(path)
		}
		return nil
	})

	// 1. Template for Constructors
	t1, err := template.New("constructors").Parse(constructorsTemplate)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if f, err := os.OpenFile(fmt.Sprintf("%s/gen_constructors.go", path), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0750); err != nil {
		log.Println("ExecuteTemplate::", err.Error())
	} else {
		t1.Execute(f, messages)
	}

	// 2. Template for Pools
	t2, err := template.New("pools").Parse(poolTemplate)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if f, err := os.OpenFile(fmt.Sprintf("%s/gen_pools.go", path), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0750); err != nil {
		log.Println("ExecuteTemplate::", err.Error())
	} else {
		t2.Execute(f, messages)
	}

	// 3. Template for Results
	t3, err := template.New("results").Parse(resultTemplate)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if f, err := os.OpenFile(fmt.Sprintf("%s/gen_results.go", path), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0750); err != nil {
		log.Println("ExecuteTemplate::", err.Error())
	} else {
		t3.Execute(f, messages)
	}

}

func ScanFile(path string) {
	var scanner *bufio.Scanner
	if file, err := os.Open(path); err != nil {
		log.Fatal(err.Error())
	} else {
		defer file.Close()
		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "message") {
			parts := strings.Split(line, " ")
			messageName := parts[1]
			constructor := crc32.ChecksumIEEE([]byte(messageName))
			if v, ok := messages[constructor]; !ok {
				messages[constructor] = messageName
			} else {
				log.Println(v, messageName)
				log.Panic("Duplicate Constructor")
			}
		}
	}
}
