package golang

import (
	"flag"
	"go/format"

	"github.com/jonsen/gencode/schema"
)

type GolangBackend struct {
	Package string
	Unsafe  bool
}

func (gb *GolangBackend) Generate(s *schema.Schema) (string, error) {
	w := &Walker{}
	w.Unsafe = gb.Unsafe
	def, err := w.WalkSchema(s, gb.Package)
	if err != nil {
		return "", err
	}
	out, err := format.Source([]byte(def.String()))
	if err != nil {
		return def.String(), nil
	}
	return string(out), nil
}

func (gb *GolangBackend) Flags() *flag.FlagSet {
	flags := flag.NewFlagSet("Go", flag.ExitOnError)
	flags.StringVar(&gb.Package, "package", "main", "package to build the gencode system for")
	flags.BoolVar(&gb.Unsafe, "unsafe", false, "Generate faster, but unsafe code")
	return flags
}

func (gb *GolangBackend) GeneratedFilename(filename string) string {
	return filename + ".gen.go"
}

func init() {
	schema.Register("go", &GolangBackend{})
}
