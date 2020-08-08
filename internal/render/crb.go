package render

import (
	"fmt"

	"github.com/openqt/osc/internal/client"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// ClusterRoleBinding renders a K8s ClusterRoleBinding to screen.
type ClusterRoleBinding struct{}

// ColorerFunc colors a resource row.
func (ClusterRoleBinding) ColorerFunc() ColorerFunc {
	return DefaultColorer
}

// Header returns a header rbw.
func (ClusterRoleBinding) Header(string) Header {
	return Header{
		HeaderColumn{Name: "NAME"},
		HeaderColumn{Name: "CLUSTERROLE"},
		HeaderColumn{Name: "SUBJECT-KIND"},
		HeaderColumn{Name: "SUBJECTS"},
		HeaderColumn{Name: "LABELS", Wide: true},
		HeaderColumn{Name: "AGE", Time: true, Decorator: AgeDecorator},
	}
}

// Render renders a K8s resource to screen.
func (ClusterRoleBinding) Render(o interface{}, ns string, r *Row) error {
	raw, ok := o.(*unstructured.Unstructured)
	if !ok {
		return fmt.Errorf("Expected ClusterRoleBinding, but got %T", o)
	}
	var crb rbacv1.ClusterRoleBinding
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(raw.Object, &crb)
	if err != nil {
		return err
	}

	kind, ss := renderSubjects(crb.Subjects)

	r.ID = client.FQN("-", crb.ObjectMeta.Name)
	r.Fields = Fields{
		crb.Name,
		crb.RoleRef.Name,
		kind,
		ss,
		mapToStr(crb.Labels),
		toAge(crb.ObjectMeta.CreationTimestamp),
	}

	return nil
}
