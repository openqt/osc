package view_test

import (
	"testing"

	"github.com/openqt/osc/internal/client"
	"github.com/openqt/osc/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestRbacNew(t *testing.T) {
	v := view.NewRbac(client.NewGVR("rbac"))

	assert.Nil(t, v.Init(makeCtx()))
	assert.Equal(t, "Rbac", v.Name())
	assert.Equal(t, 5, len(v.Hints()))
}
