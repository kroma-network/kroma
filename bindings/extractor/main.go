package main

import (
	"errors"
	"flag"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

// This code moves the Types structure from each contract binding to a separate file
// to prevent the duplication of binding for the Types structures in the contracts.

const TypesBinding = "types.go"

func main() {
	var bindingsDir string
	var name string
	var pkg string
	flag.StringVar(&bindingsDir, "dir", "", "Directory of bindings")
	flag.StringVar(&name, "name", "", "Name of binding file")
	flag.StringVar(&pkg, "pkg", "bindings", "Go package name")
	flag.Parse()

	if len(bindingsDir) == 0 {
		log.Fatalf("must define a bindings directory (example: ../bindings)")
	}

	if len(name) == 0 {
		log.Fatalf("must define a binding filename (example: helloworld.go)")
	}

	regex := regexp.MustCompile(`(?s)(?:\/\/[^\n]*\n|\/\*.*?\*\/)*\s*type\s+Types\w+\s+struct\s*\{.*?\}`)
	source, err := os.ReadFile(path.Join(bindingsDir, name))
	if err != nil {
		panic(err)
	}

	typesSource, err := os.ReadFile(path.Join(bindingsDir, TypesBinding))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			typesSource = []byte{}
		} else {
			panic(err)
		}
	}

	if len(typesSource) == 0 {
		typesSource = []byte(strings.Replace(tmpl, "{{.Package}}", pkg, 1))
	}

	for _, block := range regex.FindAllString(string(source), -1) {
		source = []byte(strings.Replace(string(source), block+"\n\n", "", -1))

		if !strings.Contains(string(typesSource), block) {
			typesSource = []byte(string(typesSource) + block + "\n\n")
		}
	}

	// formatting types bindings and save
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", typesSource, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	format.Node(io.Discard, fset, node)
	err = os.WriteFile(path.Join(bindingsDir, TypesBinding), typesSource, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// replace original code
	err = os.WriteFile(path.Join(bindingsDir, name), source, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

var tmpl = `// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.Package}}

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

`
