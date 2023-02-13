// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/substrait-io/substrait-go/extensions"
	"github.com/substrait-io/substrait-go/proto"
	"github.com/substrait-io/substrait-go/types"
)

type Relation = *proto.PlanRel

type Plan struct {
	Version            *types.Version
	Extensions         extensions.Set
	ExpectedTypeURLs   []string
	AdvancedExtensions *extensions.AdvancedExtension
	Relations          []Relation
}

func FromProto(plan *proto.Plan) (*Plan, error) {
	return &Plan{
		Version:            plan.Version,
		Relations:          plan.Relations,
		Extensions:         extensions.GetExtensionSet(plan),
		AdvancedExtensions: plan.AdvancedExtensions,
		ExpectedTypeURLs:   plan.ExpectedTypeUrls,
	}, nil
}

func (p *Plan) ToProto() (*proto.Plan, error) {
	uris, decls := p.Extensions.ToProto()
	return &proto.Plan{
		Version:            p.Version,
		ExpectedTypeUrls:   p.ExpectedTypeURLs,
		AdvancedExtensions: p.AdvancedExtensions,
		Relations:          p.Relations,
		Extensions:         decls,
		ExtensionUris:      uris,
	}, nil
}
