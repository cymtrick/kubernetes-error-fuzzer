// Code generated by pluginator on ReplicaCountTransformer; DO NOT EDIT.
// pluginator {unknown  1970-01-01T00:00:00Z  }

package builtins

import (
	"fmt"

	"sigs.k8s.io/kustomize/api/filters/replicacount"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/kustomize/kyaml/resid"
	"sigs.k8s.io/yaml"
)

// Find matching replicas declarations and replace the count.
// Eases the kustomization configuration of replica changes.
type ReplicaCountTransformerPlugin struct {
	Replica    types.Replica     `json:"replica,omitempty" yaml:"replica,omitempty"`
	FieldSpecs []types.FieldSpec `json:"fieldSpecs,omitempty" yaml:"fieldSpecs,omitempty"`
}

func (p *ReplicaCountTransformerPlugin) Config(
	_ *resmap.PluginHelpers, c []byte) (err error) {
	p.Replica = types.Replica{}
	p.FieldSpecs = nil
	return yaml.Unmarshal(c, p)
}

func (p *ReplicaCountTransformerPlugin) Transform(m resmap.ResMap) error {
	found := false
	for _, fs := range p.FieldSpecs {
		matcher := p.createMatcher(fs)
		resList := m.GetMatchingResourcesByAnyId(matcher)
		if len(resList) > 0 {
			found = true
			for _, r := range resList {
				// There are redundant checks in the filter
				// that we'll live with until resolution of
				// https://github.com/kubernetes-sigs/kustomize/issues/2506
				err := r.ApplyFilter(replicacount.Filter{
					Replica:   p.Replica,
					FieldSpec: fs,
				})
				if err != nil {
					return err
				}
			}
		}
	}

	if !found {
		gvks := make([]string, len(p.FieldSpecs))
		for i, replicaSpec := range p.FieldSpecs {
			gvks[i] = replicaSpec.Gvk.String()
		}
		return fmt.Errorf("resource with name %s does not match a config with the following GVK %v",
			p.Replica.Name, gvks)
	}

	return nil
}

// Match Replica.Name and FieldSpec
func (p *ReplicaCountTransformerPlugin) createMatcher(fs types.FieldSpec) resmap.IdMatcher {
	return func(r resid.ResId) bool {
		return r.Name == p.Replica.Name && r.Gvk.IsSelected(&fs.Gvk)
	}
}

func NewReplicaCountTransformerPlugin() resmap.TransformerPlugin {
	return &ReplicaCountTransformerPlugin{}
}
