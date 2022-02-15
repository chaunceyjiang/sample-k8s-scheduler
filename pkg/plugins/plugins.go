package plugins

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// Name 插件名称
const Name = "sample-scheduler"

type Sample struct {
	handle framework.Handle
}

// PreFilterPlugin is an interface that must be implemented by "prefilter" plugins
var _ framework.PreFilterPlugin = &Sample{}

func (s *Sample) PreFilterExtensions() framework.PreFilterExtensions {
	klog.V(3).Infof("PreFilterExtensions....")
	// PreFilterExtensions returns a PreFilterExtensions interface if the plugin implements one,
	// or nil if it does not. A Pre-filter plugin can provide extensions to incrementally
	// modify its pre-processed info. The framework guarantees that the extensions
	// AddPod/RemovePod will only be called after PreFilter, possibly on a cloned
	// CycleState, and may call those functions more than once before calling
	// Filter again on a specific node.
	return nil
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
