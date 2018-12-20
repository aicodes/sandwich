package taxonomy

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type node struct {
	Label    string
	Children map[string]*node
}

func (n *node) hasChild(child string) bool {
	_, ok := n.Children[child]
	return ok
}

//Taxonomy is a tree data structure to represent a hierarcial taxonomy
type Taxonomy struct {
	Root *node
}

func newNode(l string) *node {
	n := new(node)
	n.Label = l
	n.Children = make(map[string]*node)
	return n
}

//NewTaxonomy creates a new/empty taxonomy
func NewTaxonomy() *Taxonomy {
	return &Taxonomy{newNode("")}
}

const (
	keySep   = "." // keys are separated by "."
	valueSep = "/" // values are separated by "/"
)

//Add a new path to the taxonomy.
func (tax *Taxonomy) Add(path string) {
	levels := strings.Split(path, valueSep)
	var cur = tax.Root
	for _, v := range levels {
		if !cur.hasChild(v) {
			cur.Children[v] = newNode(v)
		}
		cur = cur.Children[v]
	}
}

//Print taxonomy in a pretty way
func (tax *Taxonomy) Print(w io.Writer) {
	cur := tax.Root
	for _, child := range cur.Children {
		fmt.Fprintf(w, "%v\n", child)
	}
}

//Init the taxonomy. Returns the number of elements loaded.
func (tax *Taxonomy) Init(fname string) (int, error) {
	values := make(map[string]string)

	jf, err := os.Open(fname)
	if err != nil {
		return 0, err
	}
	defer jf.Close()

	byteValue, _ := ioutil.ReadAll(jf)

	json.Unmarshal([]byte(byteValue), &values)

	count := 0
	for _, v := range values {
		tax.Add(v)
		count++
	}
	return count, nil
}
