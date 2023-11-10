package extractor

import (
	"errors"
	"fmt"
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

func ExtractTypes(outDir string, name string, pkg string) {
	if len(outDir) == 0 {
		log.Fatalf("must define a bindings directory (example: ../bindings)")
	}

	if len(name) == 0 {
		log.Fatalf("must define a binding filename (example: helloworld.go)")
	}

	filename := strings.ToLower(name) + ".go"
	regex := regexp.MustCompile(`(?s)(?:\/\/[^\n]*\n|\/\*.*?\*\/)*\s*type\s+Types\w+\s+struct\s*\{.*?\}`)
	source, err := os.ReadFile(path.Join(outDir, filename))
	if err != nil {
		panic(err)
	}

	output, err := os.ReadFile(path.Join(outDir, TypesBinding))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			output = []byte{}
		} else {
			panic(err)
		}
	}

	if len(output) == 0 {
		output = []byte(strings.Replace(typesTmpl, "{{.Package}}", pkg, 1))
	}

	for _, block := range regex.FindAllString(string(source), -1) {
		// remove code block from binding file.
		source = []byte(strings.Replace(string(source), block+"\n\n", "", -1))

		// find legacy code block corresponding to a given block type.
		structName := regexp.MustCompile(`Types\w+`).FindString(block)
		legacyFindRegex := regexp.MustCompile(fmt.Sprintf(
			`(?s)(?:\/\/[^\n]*\n|\/\*.*?\*\/)*\s*type\s+%s+\s+struct\s*\{.*?\}`,
			structName,
		))
		legacyBlock := legacyFindRegex.FindString(string(output))

		if len(legacyBlock) == 0 {
			// append code block if there is no legacy code block.
			output = []byte(string(output) + block + "\n\n")
		} else if legacyBlock != block {
			// else, replace if there is no legacy code block.
			output = []byte(strings.Replace(string(output), legacyBlock, block, -1))
		}
	}

	// formatting types bindings and save
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", output, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	format.Node(io.Discard, fset, node)
	err = os.WriteFile(path.Join(outDir, TypesBinding), output, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// replace original code
	err = os.WriteFile(path.Join(outDir, filename), source, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

var typesTmpl = `// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.Package}}

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

`
