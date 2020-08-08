package view_test

import (
	"testing"

	"github.com/openqt/osc/internal/client"
	"github.com/openqt/osc/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestConfigMapNew(t *testing.T) {
	s := view.NewConfigMap(client.NewGVR("v1/configmaps"))

	assert.Nil(t, s.Init(makeCtx()))
	assert.Equal(t, "ConfigMaps", s.Name())
	assert.Equal(t, 6, len(s.Hints()))
}
