package plugins

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// Name 插件名称
const Name = "sample-plugin"

type Sample struct {
	handle framework.Handle
}

func (s *Sample) Name() string {
	return Name
}

func (s *Sample) PreFilter(ctx context.Context, state *framework.CycleState, p *v1.Pod) *framework.Status {
	klog.V(3).Infof("prefilter pod: %v", p.Name)
	return framework.NewStatus(framework.Success, "")
}
func New(obj runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	return &Sample{handle: handle}, nil
}
