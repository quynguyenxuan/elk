package elk

import (
	"entgo.io/ent/entc/gen"
	"github.com/mailru/easyjson/bootstrap"
	"path/filepath"
)

type EasyJsonConfig struct {
	NoStdMarshalers          bool
	SnakeCase                bool
	LowerCamelCase           bool
	OmitEmpty                bool
	DisallowUnknownFields    bool
	SkipMemberNameUnescaping bool
}

func newEasyJsonConfig() EasyJsonConfig {
	return EasyJsonConfig{
		NoStdMarshalers:       true,
		DisallowUnknownFields: true,
	}
}

func GenerateEasyJSON(c EasyJsonConfig) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			// Let ent create all the files.
			if err := next.Generate(g); err != nil {
				return err
			}

			// We want to render every response struct created with easyjson.
			var ns []string
			vs, err := responseViews(g)
			if err != nil {
				return err
			}
			for _, v := range vs {
				n, err := v.ViewName()
				if err != nil {
					return err
				}
				ns = append(ns, n, n+"s")
			}
			for _, n := range g.Nodes {
				// Add the request structs used to deserialize request bodies.
				ns = append(ns,
					n.Name+"CreateRequest",
					n.Name+"UpdateRequest",
				)
			}

			// Run the easyjson generator.
			return (&bootstrap.Generator{
				PkgPath:                  filepath.Join(g.Package, "http"),
				PkgName:                  "http",
				Types:                    ns,
				NoStdMarshalers:          c.NoStdMarshalers,
				SnakeCase:                c.SnakeCase,
				LowerCamelCase:           c.LowerCamelCase,
				OmitEmpty:                c.OmitEmpty,
				DisallowUnknownFields:    c.DisallowUnknownFields,
				SkipMemberNameUnescaping: c.SkipMemberNameUnescaping,
				OutName:                  filepath.Join(g.Config.Target, "http", "easyjson.go"),
			}).Run()
		})
	}
}
