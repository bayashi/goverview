package fileinfo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type info struct {
	list []string
	m    map[string]struct{}
}

type organizer struct {
	structs   [2]*info
	constants [2]*info
	functions [2]*info
}

const (
	PRIV = 0
	PUB  = 1
)

func defInfo() *info {
	return &info{list: []string{}, m: map[string]struct{}{}}
}

// GoInfo provides golang file *.go info
func GoInfo(args *Args) (*FileInfo, error) {
	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, args.FilePath, nil, 0)
	if err != nil {
		return &FileInfo{}, err
	}
	//ast.Print(fset, fileAst)

	var i = organizer{
		structs:   [2]*info{defInfo(), defInfo()},
		constants: [2]*info{defInfo(), defInfo()},
		functions: [2]*info{defInfo(), defInfo()},
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

	descr := []string{}
	if !args.HideTest || !strings.HasSuffix(args.FilePath, "_test.go") {
		descr = *buildDescriptions(&i, args)
	}

	return &FileInfo{
		Icon:         "*",
		Tag:          fileAst.Name.String(),
		Descriptions: descr,
	}, nil
}

func idx(isExported bool) int {
	if isExported {
		return PUB
	}

	return PRIV
}

func goInfoProcessIdent(o *ast.Object, i *organizer) {
	switch o.Kind {
	case ast.Fun:
		f := i.functions[idx(ast.IsExported(o.Name))]
		if _, isExists := f.m[o.Name]; !isExists {
			f.list = append(f.list, o.Name)
			f.m[o.Name] = struct{}{}
		}
	case ast.Typ:
		s := i.structs[idx(ast.IsExported(o.Name))]
		if _, isExists := s.m[o.Name]; !isExists {
			s.list = append(s.list, o.Name)
			s.m[o.Name] = struct{}{}
		}
	case ast.Con:
		c := i.constants[idx(ast.IsExported(o.Name))]
		if _, isExists := c.m[o.Name]; !isExists {
			c.list = append(c.list, o.Name)
			c.m[o.Name] = struct{}{}
		}
	}
}

func buildDescriptions(i *organizer, args *Args) *[]string {
	descriptions := []string{}

	if len(i.structs[PUB].list) > 0 {
		descriptions = append(descriptions, "Struct: "+strings.Join(i.structs[PUB].list, ", "))
	}
	if args.ShowAll && len(i.structs[PRIV].list) > 0 {
		descriptions = append(descriptions, "struct: "+strings.Join(i.structs[PRIV].list, ", "))
	}

	if len(i.functions[PUB].list) > 0 {
		descriptions = append(descriptions, "Func: "+strings.Join(i.functions[PUB].list, ", "))
	}
	if args.ShowAll && len(i.functions[PRIV].list) > 0 {
		descriptions = append(descriptions, "func: "+strings.Join(i.functions[PRIV].list, ", "))
	}

	if len(i.constants[PUB].list) > 0 {
		descriptions = append(descriptions, "Const: "+strings.Join(i.constants[PUB].list, ", "))
	}
	if args.ShowAll && len(i.constants[PRIV].list) > 0 {
		descriptions = append(descriptions, "const: "+strings.Join(i.constants[PRIV].list, ", "))
	}

	return &descriptions
}
