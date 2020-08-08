package dao_test

import (
	"context"
	"testing"

	"github.com/openqt/osc/internal"
	"github.com/openqt/osc/internal/client"
	"github.com/openqt/osc/internal/dao"
	"github.com/openqt/osc/internal/render"
	"github.com/stretchr/testify/assert"
)

func TestBenchmarkList(t *testing.T) {
	a := dao.Benchmark{}
	a.Init(makeFactory(), client.NewGVR("benchmarks"))

	ctx := context.WithValue(context.Background(), internal.KeyDir, "testdata/bench")
	oo, err := a.List(ctx, "-")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(oo))
	assert.Equal(t, "testdata/bench/default_fred_1577308050814961000.txt", oo[0].(render.BenchInfo).Path)
}
