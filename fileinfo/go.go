package fileinfo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type organizer struct {
	structs            []string
	structMap          map[string]bool
	privateStructs     []string
	privateStructMap   map[string]bool
	constants          []string
	constantMap        map[string]bool
	privateConstants   []string
	privateConstantMap map[string]bool
	functions          []string
	functionMap        map[string]bool
	privateFunctions   []string
	privateFunctionMap map[string]bool
}

// GoInfo provides golang file *.go info
func GoInfo(filepath string, showAll bool) (*FileInfo, error) {
	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		return &FileInfo{}, err
	}
	//ast.Print(fset, fileAst)

	var i = organizer{
		structs:            []string{},
		structMap:          map[string]bool{},
		privateStructs:     []string{},
		privateStructMap:   map[string]bool{},
		constants:          []string{},
		constantMap:        map[string]bool{},
		privateConstants:   []string{},
		privateConstantMap: map[string]bool{},
		functions:          []string{},
		functionMap:        map[string]bool{},
		privateFunctions:   []string{},
		privateFunctionMap: map[string]bool{},
	}

	ast.Inspect(fileAst, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok {
			o := ident.Obj
			if o == nil {
				return true
			}
			goInfoProcessIdent(o, &i)
		}

		return true
	})

	return &FileInfo{
		Icon:         "*",
		Tag:          fileAst.Name.String(),
		Descriptions: *buildDescriptions(&i, showAll),
	}, nil
}

func goInfoProcessIdent(o *ast.Object, i *organizer) {
	switch o.Kind {
	case ast.Fun:
		if ast.IsExported(o.Name) {
			if _, isExists := i.functionMap[o.Name]; !isExists {
				i.functions = append(i.functions, o.Name)
				i.functionMap[o.Name] = true
			}
		} else {
			if _, isExists := i.privateFunctionMap[o.Name]; !isExists {
				i.privateFunctions = append(i.privateFunctions, o.Name)
				i.privateFunctionMap[o.Name] = true
			}
		}
	case ast.Typ:
		if ast.IsExported(o.Name) {
			if _, isExists := i.structMap[o.Name]; !isExists {
				i.structs = append(i.structs, o.Name)
				i.structMap[o.Name] = true
			}
		} else {
			if _, isExists := i.privateStructMap[o.Name]; !isExists {
				i.privateStructs = append(i.privateStructs, o.Name)
				i.privateStructMap[o.Name] = true
			}
		}
	case ast.Con:
		if ast.IsExported(o.Name) {
			if _, isExists := i.constantMap[o.Name]; !isExists {
				i.constants = append(i.constants, o.Name)
				i.constantMap[o.Name] = true
			}
		} else {
			if _, isExists := i.privateConstantMap[o.Name]; !isExists {
				i.privateConstants = append(i.privateConstants, o.Name)
				i.privateConstantMap[o.Name] = true
			}
		}
	}
}

func buildDescriptions(i *organizer, showAll bool) *[]string {
	descriptions := []string{}
	if len(i.structs) > 0 {
		descriptions = append(descriptions, "Struct: "+strings.Join(i.structs, ", "))
	}
	if len(i.functions) > 0 {
		descriptions = append(descriptions, "Func: "+strings.Join(i.functions, ", "))
	}
	if len(i.constants) > 0 {
		descriptions = append(descriptions, "Const: "+strings.Join(i.constants, ", "))
	}

	if showAll {
		if len(i.privateStructs) > 0 {
			descriptions = append(descriptions, "struct: "+strings.Join(i.privateStructs, ", "))
		}
		if len(i.privateFunctions) > 0 {
			descriptions = append(descriptions, "func: "+strings.Join(i.privateFunctions, ", "))
		}
		if len(i.privateConstants) > 0 {
			descriptions = append(descriptions, "const: "+strings.Join(i.privateConstants, ", "))
		}
	}

	return &descriptions
}
