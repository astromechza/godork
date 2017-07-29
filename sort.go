package godork

import (
	"strings"
)

type examplesByName []*Example

func (a examplesByName) Len() int           { return len(a) }
func (a examplesByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a examplesByName) Less(i, j int) bool { return strings.Compare(a[i].Name, a[j].Name) < 0 }

type functionsByName []*Function

func (a functionsByName) Len() int           { return len(a) }
func (a functionsByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a functionsByName) Less(i, j int) bool { return strings.Compare(a[i].Name, a[j].Name) < 0 }

type typesByName []*Type

func (a typesByName) Len() int           { return len(a) }
func (a typesByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a typesByName) Less(i, j int) bool { return strings.Compare(a[i].Name, a[j].Name) < 0 }

type methodsByName []*Method

func (a methodsByName) Len() int           { return len(a) }
func (a methodsByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a methodsByName) Less(i, j int) bool { return strings.Compare(a[i].Name, a[j].Name) < 0 }
