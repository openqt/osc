package view_test

import (
	"testing"

	"github.com/openqt/osc/internal/client"
	"github.com/openqt/osc/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestReferenceNew(t *testing.T) {
	s := view.NewReference(client.NewGVR("references"))

	assert.Nil(t, s.Init(makeCtx()))
	assert.Equal(t, "References", s.Name())
	assert.Equal(t, 4, len(s.Hints()))
}
