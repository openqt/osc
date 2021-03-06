package view_test

import (
	"testing"

	"github.com/openqt/osc/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestDir(t *testing.T) {
	v := view.NewDir("/fred")

	assert.Nil(t, v.Init(makeCtx()))
	assert.Equal(t, "Directory", v.Name())
	assert.Equal(t, 7, len(v.Hints()))
}
