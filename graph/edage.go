package graph

import (
	"fmt"
)

type Edage struct {
	Start *Node
	Dest  *Node
}

var EdageMap = map[*Edage]int{}

func (e *Edage) Dot(m map[*Edage]int) string {
	statics := make(map[string]int)
	for k, v := range m {
		statics[k.Start.Key()+k.Dest.Key()] += v
	}
	//taillabel
	return fmt.Sprintf("\t%s:%s -> %s: %s [label= %d ];\n", e.Start.Key(), e.Start.getType(), e.Dest.Key(), e.Dest.Key(), statics[e.Start.Key()+e.Dest.Key()])
}

func NewEdage(start, dest *Node) *Edage {
	return &Edage{
		Start: start,
		Dest:  dest,
	}
}

func AddEdage(nos, noe *Node) {
	EdageMap[NewEdage(nos, noe)]++
}
