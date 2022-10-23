// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package model

import (
	"github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/pb33f/libopenapi/what-changed/core"
)

// InfoChanges represents the number of changes to an Info object. Part of an OpenAPI document
type InfoChanges struct {
	core.PropertyChanges
	ContactChanges *ContactChanges
	LicenseChanges *LicenseChanges
}

// TotalChanges represents the total number of changes made to an Info object.
func (i *InfoChanges) TotalChanges() int {
	t := i.PropertyChanges.TotalChanges()
	if i.ContactChanges != nil {
		t += i.ContactChanges.TotalChanges()
	}
	if i.LicenseChanges != nil {
		t += i.LicenseChanges.TotalChanges()
	}
	return t
}

// TotalBreakingChanges always returns 0 for Info objects, they are non-binding.
func (i *InfoChanges) TotalBreakingChanges() int {
	return 0
}

// CompareInfo will compare a left (original) and a right (new) Info object. Any changes
// will be returned in a pointer to InfoChanges, otherwise if nothing is found, then nil is
// returned instead.
func CompareInfo(l, r *base.Info) *InfoChanges {

	var changes []*core.Change
	var props []*core.PropertyCheck

	// Title
	props = append(props, &core.PropertyCheck{
		LeftNode:  l.Title.ValueNode,
		RightNode: r.Title.ValueNode,
		Label:     v3.TitleLabel,
		Changes:   &changes,
		Breaking:  false,
		Original:  l,
		New:       r,
	})

	// Description
	props = append(props, &core.PropertyCheck{
		LeftNode:  l.Description.ValueNode,
		RightNode: r.Description.ValueNode,
		Label:     v3.DescriptionLabel,
		Changes:   &changes,
		Breaking:  false,
		Original:  l,
		New:       r,
	})

	// TermsOfService
	props = append(props, &core.PropertyCheck{
		LeftNode:  l.TermsOfService.ValueNode,
		RightNode: r.TermsOfService.ValueNode,
		Label:     v3.TermsOfServiceLabel,
		Changes:   &changes,
		Breaking:  false,
		Original:  l,
		New:       r,
	})

	// Version
	props = append(props, &core.PropertyCheck{
		LeftNode:  l.Version.ValueNode,
		RightNode: r.Version.ValueNode,
		Label:     v3.VersionLabel,
		Changes:   &changes,
		Breaking:  false,
		Original:  l,
		New:       r,
	})

	// check properties
	core.CheckProperties(props)

	i := new(InfoChanges)

	// compare contact.
	if l.Contact.Value != nil && r.Contact.Value != nil {
		i.ContactChanges = CompareContact(l.Contact.Value, r.Contact.Value)
	} else {
		if l.Contact.Value == nil && r.Contact.Value != nil {
			core.CreateChange(&changes, core.ObjectAdded, v3.ContactLabel,
				nil, r.Contact.ValueNode, false, nil, r.Contact.Value)
		}
		if l.Contact.Value != nil && r.Contact.Value == nil {
			core.CreateChange(&changes, core.ObjectRemoved, v3.ContactLabel,
				l.Contact.ValueNode, nil, false, l.Contact.Value, nil)
		}
	}

	// compare license.
	if l.License.Value != nil && r.License.Value != nil {
		i.LicenseChanges = CompareLicense(l.License.Value, r.License.Value)
	} else {
		if l.License.Value == nil && r.License.Value != nil {
			core.CreateChange(&changes, core.ObjectAdded, v3.LicenseLabel,
				nil, r.License.ValueNode, false, nil, r.License.Value)
		}
		if l.License.Value != nil && r.License.Value == nil {
			core.CreateChange(&changes, core.ObjectRemoved, v3.LicenseLabel,
				l.License.ValueNode, nil, false, r.License.Value, nil)
		}
	}
	i.Changes = changes
	if i.TotalChanges() <= 0 {
		return nil
	}
	return i
}