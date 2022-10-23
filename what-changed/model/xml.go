// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low/base"
	v3 "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/what-changed/core"
)

// XMLChanges represents changes made to the XML object of an OpenAPI document.
type XMLChanges struct {
	core.PropertyChanges
	ExtensionChanges *ExtensionChanges
}

// TotalChanges returns a count of everything that was changed within an XML object.
func (x *XMLChanges) TotalChanges() int {
	c := x.PropertyChanges.TotalChanges()
	if x.ExtensionChanges != nil {
		c += x.ExtensionChanges.TotalChanges()
	}
	return c
}

// TotalBreakingChanges returns the number of breaking changes made by the XML object.
func (x *XMLChanges) TotalBreakingChanges() int {
	return x.PropertyChanges.TotalBreakingChanges()
}

// CompareXML will compare a left (original) and a right (new) XML instance, and check for
// any changes between them. If changes are found, the function returns a pointer to XMLChanges,
// otherwise, if nothing changed - it will return nil
func CompareXML(l, r *base.XML) *XMLChanges {
	xc := new(XMLChanges)
	var changes []*core.Change
	var props []*core.PropertyCheck

	// Name (breaking change)
	props = append(props, &core.PropertyCheck{
		LeftNode:  l.Name.ValueNode,
		RightNode: r.Name.ValueNode,
		Label:     v3.NameLabel,
		Changes:   &changes,
		Breaking:  true,
		Original:  l,
		New:       r,
	})

	// Namespace (breaking change)
	props = append(props, &core.PropertyCheck{
		LeftNode:  l.Namespace.ValueNode,
		RightNode: r.Namespace.ValueNode,
		Label:     v3.NamespaceLabel,
		Changes:   &changes,
		Breaking:  true,
		Original:  l,
		New:       r,
	})

	// Prefix (breaking change)
	props = append(props, &core.PropertyCheck{
		LeftNode:  l.Prefix.ValueNode,
		RightNode: r.Prefix.ValueNode,
		Label:     v3.PrefixLabel,
		Changes:   &changes,
		Breaking:  true,
		Original:  l,
		New:       r,
	})

	// Attribute (breaking change)
	props = append(props, &core.PropertyCheck{
		LeftNode:  l.Attribute.ValueNode,
		RightNode: r.Attribute.ValueNode,
		Label:     v3.AttributeLabel,
		Changes:   &changes,
		Breaking:  true,
		Original:  l,
		New:       r,
	})

	// Wrapped (breaking change)
	props = append(props, &core.PropertyCheck{
		LeftNode:  l.Wrapped.ValueNode,
		RightNode: r.Wrapped.ValueNode,
		Label:     v3.WrappedLabel,
		Changes:   &changes,
		Breaking:  true,
		Original:  l,
		New:       r,
	})

	// check properties
	core.CheckProperties(props)

	// check extensions
	xc.ExtensionChanges = CheckExtensions(l, r)
	xc.Changes = changes
	if xc.TotalChanges() <= 0 {
		return nil
	}
	return xc
}