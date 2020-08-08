package view_test

import (
	"testing"

	"github.com/openqt/osc/internal/client"
	"github.com/openqt/osc/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestPVCNew(t *testing.T) {
	v := view.NewPersistentVolumeClaim(client.NewGVR("v1/persistentvolumeclaims"))

	assert.Nil(t, v.Init(makeCtx()))
	assert.Equal(t, "PersistentVolumeClaims", v.Name())
	assert.Equal(t, 10, len(v.Hints()))
}
