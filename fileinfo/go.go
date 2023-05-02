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

type p struct {
	public  *info
	private *info
}

type organizer struct {
	structs   *p
	constants *p
	functions *p
}

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
		structs:   &p{public: defInfo(), private: defInfo()},
		constants: &p{public: defInfo(), private: defInfo()},
		functions: &p{public: defInfo(), private: defInfo()},
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

func goInfoProcessIdent(o *ast.Object, i *organizer) {
	switch o.Kind {
	case ast.Fun:
		if ast.IsExported(o.Name) {
			if _, isExists := i.functions.public.m[o.Name]; !isExists {
				i.functions.public.list = append(i.functions.public.list, o.Name)
				i.functions.public.m[o.Name] = struct{}{}
			}
		} else {
			if _, isExists := i.functions.private.m[o.Name]; !isExists {
				i.functions.private.list = append(i.functions.private.list, o.Name)
				i.functions.private.m[o.Name] = struct{}{}
			}
		}
	case ast.Typ:
		if ast.IsExported(o.Name) {
			if _, isExists := i.structs.public.m[o.Name]; !isExists {
				i.structs.public.list = append(i.structs.public.list, o.Name)
				i.structs.public.m[o.Name] = struct{}{}
			}
		} else {
			if _, isExists := i.structs.private.m[o.Name]; !isExists {
				i.structs.private.list = append(i.structs.private.list, o.Name)
				i.structs.private.m[o.Name] = struct{}{}
			}
		}
	case ast.Con:
		if ast.IsExported(o.Name) {
			if _, isExists := i.constants.public.m[o.Name]; !isExists {
				i.constants.public.list = append(i.constants.public.list, o.Name)
				i.constants.public.m[o.Name] = struct{}{}
			}
		} else {
			if _, isExists := i.constants.private.m[o.Name]; !isExists {
				i.constants.private.list = append(i.constants.private.list, o.Name)
				i.constants.private.m[o.Name] = struct{}{}
			}
		}
	}
}

func buildDescriptions(i *organizer, args *Args) *[]string {
	descriptions := []string{}
	if len(i.structs.public.list) > 0 {
		descriptions = append(descriptions, "Struct: "+strings.Join(i.structs.public.list, ", "))
	}
	if args.ShowAll && len(i.structs.private.list) > 0 {
		descriptions = append(descriptions, "struct: "+strings.Join(i.structs.private.list, ", "))
	}
	if len(i.functions.public.list) > 0 {
		descriptions = append(descriptions, "Func: "+strings.Join(i.functions.public.list, ", "))
	}
	if args.ShowAll && len(i.functions.private.list) > 0 {
		descriptions = append(descriptions, "func: "+strings.Join(i.functions.private.list, ", "))
	}
	if len(i.constants.public.list) > 0 {
		descriptions = append(descriptions, "Const: "+strings.Join(i.constants.public.list, ", "))
	}
	if args.ShowAll && len(i.constants.private.list) > 0 {
		descriptions = append(descriptions, "const: "+strings.Join(i.constants.private.list, ", "))
	}

	return &descriptions
}
