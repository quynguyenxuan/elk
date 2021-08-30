package elk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template/parse"

	"strings"
	"unicode"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/masseelch/elk/internal"

	"golang.org/x/tools/imports"
)

type (
	// Extension implements entc.Extension interface for providing http handler code generation.
	Extension struct {
		entc.DefaultExtension
		easyjsonConfig EasyJsonConfig
		hooks          []gen.Hook
		templates      []*gen.Template
	}
	// ExtensionOption allows to manage Extension configuration using functional arguments.
	ExtensionOption func(*Extension) error
)

type (
	file struct {
		path    string
		content []byte
	}
	assets struct {
		dirs  []string
		files []file
	}
	Graph struct {
		gen.Graph
	}
)

// NewExtension returns a new elk extension with default values.
func NewExtension(opts ...ExtensionOption) (*Extension, error) {
	ex := &Extension{
		templates:      HTTPTemplates,
		easyjsonConfig: newEasyJsonConfig(),
	}
	for _, opt := range opts {
		if err := opt(ex); err != nil {
			return nil, err
		}
	}
	ex.hooks = append(ex.hooks, GenerateEasyJSON(ex.easyjsonConfig))
	// ex.hooks = append(ex.hooks, GenerateHttpHandler())
	return ex, nil
}

func (e *Extension) Templates() []*gen.Template {
	return e.templates
}

func (e *Extension) Hooks() []gen.Hook {
	return e.hooks
}

func WithEasyJsonConfig(c EasyJsonConfig) ExtensionOption {
	return func(ex *Extension) error {
		ex.easyjsonConfig = c
		return nil
	}
}

func GenerateHttpHandler() gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			var (
				assets assets
				// external []gen.GraphTemplate
			)
			templates, _ = gTemplates(g)
			// Templates = append(Templates, )
			for _, n := range g.Nodes {
				assets.dirs = append(assets.dirs, filepath.Join(g.Config.Target, n.Package()))
				for _, tmpl := range ExtendHttpTemplates {
					b := bytes.NewBuffer(nil)
					if err := templates.ExecuteTemplate(b, tmpl.Name, n); err != nil {
						return fmt.Errorf("execute template %q: %w", tmpl.Name, err)
					}
					assets.files = append(assets.files, file{
						path:    filepath.Join(g.Config.Target, tmpl.Format(n)),
						content: b.Bytes(),
					})
				}
			}

			// Write and format assets only if template execution
			// finished successfully.
			if err := assets.write(); err != nil {
				return err
			}
			// We can't run "imports" on files when the state is not completed.
			// Because, "goimports" will drop undefined package. Therefore, it's
			// suspended to the end of the writing.
			return assets.format()
		})
	}
}

// templates returns the template.Template for the code and external templates
// to execute on the Graph object if provided.
func gTemplates(g *gen.Graph) (*gen.Template, []gen.GraphTemplate) {
	initTemplates()
	external := make([]gen.GraphTemplate, 0, len(g.Templates))
	for _, rootT := range g.Templates {
		templates.Funcs(TemplateFuncs)
		for _, tmpl := range rootT.Templates() {
			if parse.IsEmptyTree(tmpl.Root) {
				continue
			}
			name := tmpl.Name()
			// If the template does not override or extend one of
			// the builtin templates, generate it in a new file.
			// if templates.Lookup(name) == nil && !extendExisting(name) {
			// 	external = append(external, gen.GraphTemplate{
			// 		Name:   name,
			// 		Format: snake(name) + ".go",
			// 	})
			// }
			templates = gen.MustParse(templates.AddParseTree(name, tmpl.Tree))
		}
	}
	// for _, f := range g.Features {
	// 	external = append(external, f.GraphTemplates...)
	// }
	return templates, external
}

// snake converts the given struct or field name into a snake_case.
//
//	Username => username
//	FullName => full_name
//	HTTPCode => http_code
//
func snake(s string) string {
	var (
		j int
		b strings.Builder
	)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		// Put '_' if it is not a start or end of a word, current letter is uppercase,
		// and previous is lowercase (cases like: "UserInfo"), or next letter is also
		// a lowercase and previous letter is not "_".
		if i > 0 && i < len(s)-1 && unicode.IsUpper(r) {
			if unicode.IsLower(rune(s[i-1])) ||
				j != i-1 && unicode.IsLower(rune(s[i+1])) && unicode.IsLetter(rune(s[i-1])) {
				j = i
				b.WriteString("_")
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

func match(patterns []string, name string) bool {
	for _, pat := range patterns {
		matched, _ := filepath.Match(pat, name)
		if matched {
			return true
		}
	}
	return false
}

func extendExisting(name string) bool {
	for _, t := range ExtendHttpTemplates {
		if match(t.ExtendPatterns, name) {
			return true
		}
	}
	return false
}

// write files and dirs in the assets.
func (a assets) write() error {
	for _, dir := range a.dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("create dir %q: %w", dir, err)
		}
	}
	for _, file := range a.files {
		if err := ioutil.WriteFile(file.path, file.content, 0644); err != nil {
			return fmt.Errorf("write file %q: %w", file.path, err)
		}
	}
	return nil
}

// format runs "goimports" on all assets.
func (a assets) format() error {
	for _, file := range a.files {
		path := file.path
		src, err := imports.Process(path, file.content, nil)
		if err != nil {
			return fmt.Errorf("format file %s: %w", path, err)
		}
		if err := ioutil.WriteFile(path, src, 0644); err != nil {
			return fmt.Errorf("write file %s: %w", path, err)
		}
	}
	return nil
}

func initTemplates() {
	templates = gen.NewTemplate("templates")

	templates.Funcs(TemplateFuncs)
	fmt.Println(internal.AssetNames())
	for _, asset := range internal.AssetNames() {
		templates = gen.MustParse(templates.Parse(string(internal.MustAsset(asset))))
		fmt.Println(asset)
	}
	// b := bytes.NewBuffer([]byte("package main\n"))
	// check(templates.ExecuteTemplate(b, "import", gen.Type{Config: &gen.Config{}}), "load imports")
	// f, err := parser.ParseFile(token.NewFileSet(), "", b, parser.ImportsOnly)
	// check(err, "parse imports")
	// for _, spec := range f.Imports {
	// 	path, err := strconv.Unquote(spec.Path.Value)
	// 	check(err, "unquote import path")
	// 	importPkg[filepath.Base(path)] = path
	// }
	// for _, s := range drivers {
	// 	for _, path := range s.Imports {
	// 		importPkg[filepath.Base(path)] = path
	// 	}
	// }
}

// check panics if the error is not nil.
func check(err error, msg string, args ...interface{}) {
	if err != nil {
		args = append(args, err)
		panic(graphError{fmt.Sprintf(msg+": %s", args...)})
	}
}
