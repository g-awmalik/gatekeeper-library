package suite

import (
	"fmt"
	"gatekeeperlibrary/cmd/pkg/apis"
	"os"
	"path"
	"sort"

	gkapis "github.com/open-policy-agent/gatekeeper/apis"
	"github.com/open-policy-agent/gatekeeper/pkg/gktest"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const Strategy string = "suite"

var scheme *runtime.Scheme

func init() {
	scheme = runtime.NewScheme()
	err := gkapis.AddToScheme(scheme)
	if err != nil {
		panic(fmt.Errorf("adding gatekeeper apis to scheme: %w", err))
	}
}

// Ingest takes an absolute path to a directory containing suite.yaml files and
// accompanying library resources.  These includes templates, constraints, and
// example objects.  It reads their data into memory and returns as a slice of
// ConstraintTemplateDoc structs.
func Ingest(root string) ([]*apis.ConstraintTemplateDoc, error) {
	system := os.DirFS(root)

	sMap, err := gktest.ReadSuites(system, ".", true)
	if err != nil {
		return nil, fmt.Errorf("reading suites: %w", err)
	}

	// Because multiple suite Tests can have the same template, we need to be
	// prepared to add data to existing ConstraintTemplateDoc structs.
	// Organizing them in a map makes this easier.
	templateMap := map[string]*apis.ConstraintTemplateDoc{}
	for _, suite := range sMap {
		suiteDir := path.Dir(suite.Path)
		for _, test := range suite.Tests {
			tmpl, err := gktest.ReadTemplate(scheme, system, path.Join(suiteDir, test.Template))
			if err != nil {
				return nil, fmt.Errorf("reading template of test %q, suite %q: %w", test.Name, suite.Name, err)
			}

			constraint, err := gktest.ReadObject(system, path.Join(suiteDir, test.Constraint))
			if err != nil {
				return nil, fmt.Errorf("reading constraint of test %q, suite %q: %w", test.Name, suite.Name, err)
			}

			//check if the constraint has the annotation to skip doc-gen for this constraint
			docGenFlag := constraint.GetAnnotations()["policy.library/doc-gen"]
			if docGenFlag == "do_not_document" {
				continue
			}

			// If we haven't added this template to the map yet, create the data structure
			if _, ok := templateMap[tmpl.GetName()]; !ok {
				templateMap[tmpl.GetName()] = &apis.ConstraintTemplateDoc{
					Name:        tmpl.Spec.CRD.Spec.Names.Kind,
					Description: tmpl.Annotations["description"],
					Validation:  tmpl.Spec.CRD.Spec.Validation,
				}
			}

			// Define a newSample for this constraint.  Each case has an allowed
			// or disallowed object that we'll add to the newSample.
			newSample := &apis.Sample{
				Constraint: &apis.ObjectWithContext{
					Object: constraint,
				},
			}

			for _, cas := range test.Cases {
				// Read in the object under test
				obj, err := gktest.ReadObject(system, path.Join(suiteDir, cas.Object))
				if err != nil {
					return nil, fmt.Errorf("reading object for case %q, test %q, suite %q: %w", cas.Name, test.Name, suite.Name, err)
				}

				owc := &apis.ObjectWithContext{
					Object: obj,
				}

				// categorize the object has an allowed or disallowed example
				if allowedCase(cas) {
					newSample.Allowed = append(newSample.Allowed, owc)
				} else {
					newSample.Disallowed = append(newSample.Disallowed, owc)
				}

				// read in the referential data if it exists
				for _, refObjPath := range cas.Inventory {
					refObj, err := gktest.ReadObject(system, path.Join(suiteDir, refObjPath))
					if err != nil {
						return nil, fmt.Errorf("reading inventory object %q for case %q, test %q, suite %q: %w", refObjPath, cas.Name, test.Name, suite.Name, err)
					}

					// need to write a test case
					owc.ReferentialData = append(owc.ReferentialData, refObj)
				}
			}

			// It's possible that multiple tests have the same
			// template/constraint combo.  If that's the case, we need to merge
			// the Sample objects for that constraint together.  Merging in
			// place prevents us from disturbing any existing ordering.
			didMerge := false
			for i, oldSample := range templateMap[tmpl.GetName()].Samples {
				newConstraint := newSample.Constraint.Object
				oldConstraint := oldSample.Constraint.Object
				if newConstraint.GroupVersionKind() != oldConstraint.GroupVersionKind() {
					continue
				}
				if newConstraint.GetName() != oldConstraint.GetName() {
					continue
				}

				// if we find a match, merge them in place.
				merged, err := apis.Merge(newSample, oldSample)
				if err != nil {
					return nil, fmt.Errorf("merging constraint in test %q, suite %q: %w", test.Name, suite.Name, err)
				}
				templateMap[tmpl.GetName()].Samples[i] = merged
				didMerge = true
			}
			if !didMerge {
				templateMap[tmpl.GetName()].Samples = append(templateMap[tmpl.GetName()].Samples, newSample)
			}
		}
	}

	return slice(templateMap), nil
}

func allowedCase(c *gktest.Case) bool {
	if len(c.Assertions) != 1 {
		return false
	}

	violations := c.Assertions[0].Violations
	if violations.Type == intstr.String {
		if violations.StrVal == "no" {
			return true
		}
	} else {
		if violations.IntVal == 0 {
			return true
		}
	}

	return false
}

func slice(in map[string]*apis.ConstraintTemplateDoc) []*apis.ConstraintTemplateDoc {
	out := make([]*apis.ConstraintTemplateDoc, 0, len(in))

	// Key loops aren't guaranteed ordered the same every time.  Sort the keys
	// proactively.  This prevents the docs from generating differently each
	// time when there are multiple samples folders.
	keys := make([]string, 0, len(in))
	for k := range in {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		out = append(out, in[k])
	}

	return out
}
