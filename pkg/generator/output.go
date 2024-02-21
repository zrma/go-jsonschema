package generator

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	"github.com/zrma/go-jsonschema/pkg/cmputil"
	"github.com/zrma/go-jsonschema/pkg/codegen"
	"github.com/zrma/go-jsonschema/pkg/schemas"
)

type output struct {
	file          *codegen.File
	declsByName   map[string]*codegen.TypeDecl
	declsBySchema map[*schemas.Type]*codegen.TypeDecl
	warner        func(string)
}

func (o *output) getDeclByEqualSchema(name string, t *schemas.Type) *codegen.TypeDecl {
	v, ok := o.declsByName[name]
	if !ok {
		o.warner(fmt.Sprintf("Name not found: %s", name))

		return nil
	}

	if cmp.Equal(v.SchemaType, t, cmputil.Opts(*v.SchemaType, *t)...) {
		return v
	}

	for count := 1; ; count++ {
		suffixed := fmt.Sprintf("%s_%d", name, count)

		sv, ok := o.declsByName[suffixed]
		if !ok {
			return nil
		}

		if cmp.Equal(sv.SchemaType, t, cmputil.Opts(*sv.SchemaType, *t)...) {
			return sv
		}
	}
}

func (o *output) isUniqueTypeName(name string) bool {
	v, ok := o.declsByName[name]

	return !ok || (ok && v.Type == nil)
}

func (o *output) uniqueTypeName(name string) string {
	v, ok := o.declsByName[name]

	if !ok || (ok && v.Type == nil) {
		return name
	}

	count := 1

	for {
		suffixed := fmt.Sprintf("%s_%d", name, count)
		if _, ok := o.declsByName[suffixed]; !ok {
			o.warner(fmt.Sprintf(
				"Multiple types map to the name %q; declaring duplicate as %q instead", name, suffixed))

			return suffixed
		}
		count++
	}
}
