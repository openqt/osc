package model

import (
	"context"

	"github.com/openqt/osc/internal/client"
	"github.com/openqt/osc/internal/dao"
	v1 "k8s.io/api/core/v1"
	mv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

type (
	// MetricsServer gather metrics information from pods and nodes.
	MetricsServer interface {
		MetricsService

		ClusterLoad(*v1.NodeList, *mv1beta1.NodeMetricsList, *client.ClusterMetrics) error
		NodesMetrics(*v1.NodeList, *mv1beta1.NodeMetricsList, client.NodesMetrics)
		PodsMetrics(*mv1beta1.PodMetricsList, client.PodsMetrics)
	}

	// MetricsService calls the metrics server for metrics info.
	MetricsService interface {
		HasMetrics() bool
		FetchNodesMetrics(ctx context.Context) (*mv1beta1.NodeMetricsList, error)
		FetchPodsMetrics(ctx context.Context, ns string) (*mv1beta1.PodMetricsList, error)
	}

	// Cluster represents a kubernetes resource.
	Cluster struct {
		factory dao.Factory
		mx      MetricsServer
	}
)

// NewCluster returns a new cluster info resource.
func NewCluster(f dao.Factory) *Cluster {
	return &Cluster{
		factory: f,
		mx:      client.DialMetrics(f.Client()),
	}
}

// Version returns the current K8s cluster version.
func (c *Cluster) Version() string {
	info, err := c.factory.Client().ServerVersion()
	if err != nil {
		return NA
	}

	return info.GitVersion
}

// ContextName returns the context name.
func (c *Cluster) ContextName() string {
	n, err := c.factory.Client().Config().CurrentContextName()
	if err != nil {
		return NA
	}
	return n
}

// ClusterName returns the cluster name.
func (c *Cluster) ClusterName() string {
	n, err := c.factory.Client().Config().CurrentClusterName()
	if err != nil {
		return NA
	}
	return n
}

// UserName returns the user name.
func (c *Cluster) UserName() string {
	n, err := c.factory.Client().Config().CurrentUserName()
	if err != nil {
		return NA
	}
	return n
}

// Metrics gathers node level metrics and compute utilization percentages.
func (c *Cluster) Metrics(ctx context.Context, mx *client.ClusterMetrics) error {
	nn, err := dao.FetchNodes(ctx, c.factory, "")
	if err != nil {
		return err
	}

	nmx, err := c.mx.FetchNodesMetrics(ctx)
	if err != nil {
		return err
	}

	return c.mx.ClusterLoad(nn, nmx, mx)
}
