package view_test

import (
	"testing"

	"github.com/openqt/osc/internal/client"
	"github.com/openqt/osc/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestDaemonSet(t *testing.T) {
	v := view.NewDaemonSet(client.NewGVR("apps/v1/daemonsets"))

	assert.Nil(t, v.Init(makeCtx()))
	assert.Equal(t, "DaemonSets", v.Name())
	assert.Equal(t, 13, len(v.Hints()))
}
