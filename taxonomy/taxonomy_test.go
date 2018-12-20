package taxonomy

import (
	"testing"
)

func TestEmptyTaxonomy(t *testing.T) {
	tax := NewTaxonomy()
	if tax.Root.Label != "" {
		t.Error("empty taxnomony root should be empty")
	}
}

func TestAdd(t *testing.T) {
	tax := NewTaxonomy()

	tax.Add("a/b/c")
	tax.Add("a/b/d")

	a, ok := tax.Root.Children["a"]
	if !ok {
		t.Error("root should have 'a' as child.")
	}
	b, ok := a.Children["b"]
	if !ok {
		t.Error("a/ should have 'b' as child.")
	}
	if len(b.Children) != 2 {
		t.Error("b/ should have two children.")
	}
}
