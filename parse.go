package godork

import (
	"bytes"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ReadPackageFile(fset *token.FileSet, sourceFile string) (*ast.File, error) {
	f, err := os.Open(sourceFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return parser.ParseFile(fset, "", f, parser.ParseComments)
}

func ReadPackageDirectory(sourceDirectory string) (*token.FileSet, *ast.Package, error) {
	fis, err := ioutil.ReadDir(sourceDirectory)
	if err != nil {
		return nil, nil, err
	}

	fset := token.NewFileSet()
	pkg := &ast.Package{
		Files: make(map[string]*ast.File),
	}

	for _, fi := range fis {
		if !fi.IsDir() && strings.HasSuffix(fi.Name(), ".go") {
			fileName := filepath.Join(sourceDirectory, fi.Name())
			pkgFile, perr := ReadPackageFile(fset, fileName)
			if perr != nil {
				return nil, nil, perr
			}
			pkg.Name = pkgFile.Name.Name
			pkg.Files[fileName] = pkgFile
		}
	}

	return fset, pkg, err
}

func CodeToString(fset *token.FileSet, node interface{}) string {
	buf := new(bytes.Buffer)
	printer.Fprint(buf, fset, node)
	return buf.String()
}

func BuildExamplesForName(examples []*doc.Example, name string, fset *token.FileSet) []*Example {
	var out []*Example
	for _, eg := range examples {
		if strings.HasPrefix(eg.Name, name) {
			namedetail := eg.Name[:len(eg.Name)-len(name)]
			if len(namedetail) == 0 || strings.HasPrefix(namedetail, "_") {
				egg := &Example{
					Name:     eg.Name,
					FullName: "Example" + eg.Name,
					Doc:      eg.Doc,
					Code:     CodeToString(fset, eg.Code),
					Output:   eg.Output,
				}
				if eg.Play != nil {
					egg.Code = CodeToString(fset, eg.Play)
				}
				out = append(out, egg)
			}
		}
	}
	if out == nil {
		return make([]*Example, 0)
	}
	sort.Sort(examplesByName(out))
	return out
}

func BuildPackageDoc(importPath string, fset *token.FileSet, pkg *ast.Package) (*PackageDoc, error) {

	ast.PackageExports(pkg)

	seen := make(map[string]bool)
	var imports []string

	var egfiles []*ast.File
	for _, f := range pkg.Files {
		egfiles = append(egfiles, f)
		for _, is := range f.Imports {
			pps := is.Path.Value
			_, ok := seen[pps]
			if !ok {
				imports = append(imports, pps)
			}
			seen[pps] = true
		}
	}
	sort.Strings(imports)

	examples := doc.Examples(egfiles...)
	docPkg := doc.New(pkg, importPath, doc.AllDecls)

	out := new(PackageDoc)
	out.PackageName = pkg.Name
	out.ImportPath = importPath
	out.Doc = docPkg.Doc
	out.Imports = imports
	out.Examples = BuildExamplesForName(examples, "", fset)

	for _, c := range docPkg.Consts {
		cc := &Constant{
			Doc:  c.Doc,
			Code: CodeToString(fset, c.Decl),
		}
		out.Constants = append(out.Constants, cc)
	}

	for _, v := range docPkg.Vars {
		vv := &Variable{
			Doc:  v.Doc,
			Code: CodeToString(fset, v.Decl),
		}
		out.Variables = append(out.Variables, vv)
	}

FuncLoop:
	for _, f := range docPkg.Funcs {

		for _, eg := range examples {
			if "Example"+eg.Name == f.Name {
				continue FuncLoop
			}
		}

		ff := &Function{
			Name:      f.Name,
			Doc:       f.Doc,
			Signature: CodeToString(fset, f.Decl),
			Examples:  BuildExamplesForName(examples, f.Name, fset),
		}
		out.Functions = append(out.Functions, ff)
	}
	sort.Sort(functionsByName(out.Functions))

	for _, t := range docPkg.Types {
		tt := &Type{
			Name:     t.Name,
			Doc:      t.Doc,
			Code:     CodeToString(fset, t.Decl),
			Examples: BuildExamplesForName(examples, t.Name, fset),
		}
		for _, m := range t.Methods {
			mm := &Method{
				Name:      m.Name,
				Doc:       m.Doc,
				Signature: CodeToString(fset, m.Decl),
				Receiver:  m.Recv,
				Examples:  BuildExamplesForName(examples, t.Name+"_"+m.Name, fset),
			}
			tt.Methods = append(tt.Methods, mm)
		}
		sort.Sort(methodsByName(tt.Methods))
		out.Types = append(out.Types, tt)
	}
	sort.Sort(typesByName(out.Types))
	return out, nil
}
