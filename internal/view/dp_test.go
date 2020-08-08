package view_test

import (
	"testing"

	"github.com/openqt/osc/internal/client"
	"github.com/openqt/osc/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestDeploy(t *testing.T) {
	v := view.NewDeploy(client.NewGVR("apps/v1/deployments"))

	assert.Nil(t, v.Init(makeCtx()))
	assert.Equal(t, "Deployments", v.Name())
	assert.Equal(t, 12, len(v.Hints()))
}
