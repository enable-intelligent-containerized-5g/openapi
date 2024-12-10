package pcm_models

import (
	"fmt"
	"strings"
)

type PrometheusQueryParams struct {
	Namespace    string `json:"namespace"`
	Pod          string `json:"pod"`
	Container    string `json:"container"`
	TargetPeriod string `json:"targetPeriod"`
	Offset       string `json:"offset"`
	Unit         string `json:"unit"`
	Instance     string `json:"instance"`
	Phase        string `json:"phase"`
}

// UploadThrough ()
func BuildTotalThroughputQuery(p *PrometheusQueryParams) string {
	query1 := fmt.Sprintf(`sum(rate(container_network_transmit_bytes_total{namespace="%s", pod="%s", container!~".*wait-.*"}[%s])) by (pod, namespace)`,
		p.Namespace, p.Pod, p.TargetPeriod)
	query2 := fmt.Sprintf(`sum(rate(container_network_receive_bytes_total{namespace="%s", pod="%s", container!~".*wait-.*"}[%s])) by (pod, namespace)`,
		p.Namespace, p.Pod, p.TargetPeriod)

	return fmt.Sprintf(`%s + %s`, query1, query2)
}

// UploadThrough ()
func BuildUploadThroughputQuery(p *PrometheusQueryParams) string {
	return fmt.Sprintf(`sum(rate(container_network_transmit_bytes_total{namespace="%s", pod="%s", container!~".*wait-.*"}[%s])) by (pod, namespace)`,
		p.Namespace, p.Pod, p.TargetPeriod)
}

// DownloadThrough
func BuildDownloadThroughputQuery(p *PrometheusQueryParams) string {
	return fmt.Sprintf(`sum(rate(container_network_receive_bytes_total{namespace="%s", pod="%s", container!~".*wait-.*"}[%s])) by (pod, namespace)`,
		p.Namespace, p.Pod, p.TargetPeriod)
}

// Memory rate (OK)
func BuildMemRateQuery(p *PrometheusQueryParams) string {
	return fmt.Sprintf(`sum(rate(container_memory_usage_bytes{namespace="%s", pod="%s", container="%s", container!~".*wait-.*"}[%s] offset %s)) by (pod, container, namespace)`,
		p.Namespace, p.Pod, p.Container, p.TargetPeriod, p.Offset)
}

// Memory Usage (OK)
func BuildMemUsageQuery(p *PrometheusQueryParams) string {
	return fmt.Sprintf(`avg(container_memory_usage_bytes{namespace="%s", pod="%s", container="%s", container!~"wait-.*"}) by (pod, container, namespace)`,
		p.Namespace, p.Pod, p.Container)
}

// CPU Usage Average (OK)
func BuildCpuUsageAverageQuery(p *PrometheusQueryParams) string {
	var offsetQuery string
	offSet := p.Offset
	if strings.TrimSpace(offSet) != "" {
		offsetQuery = fmt.Sprintf(` offset %s`, p.Offset)
	}
	return fmt.Sprintf(`avg(rate(container_cpu_usage_seconds_total{namespace="%s", pod="%s", container="%s", container!~".*wait-.*"}[%s]%s)) by (pod, container, namespace)`,
		p.Namespace, p.Pod, p.Container, p.TargetPeriod, offsetQuery)
}

// Memory Usage average (OK)
func BuildMemUsageAverageQuery(p *PrometheusQueryParams) string {
	var offsetQuery string
	offSet := p.Offset
	if strings.TrimSpace(offSet) != "" {
		offsetQuery = fmt.Sprintf(` offset %s`, p.Offset)
	}
	return fmt.Sprintf(`avg(avg_over_time(container_memory_usage_bytes{namespace="%s", pod="%s", container="%s", container!~"wait-.*"}[%s]%s)) by (pod, container, namespace)`,
		p.Namespace, p.Pod, p.Container, p.TargetPeriod, offsetQuery)
}

// CPU and Memory resources request (OK)
func BuildResourceRequestQuery(p *PrometheusQueryParams) string {
	return fmt.Sprintf(`avg(kube_pod_container_resource_requests{namespace="%s", pod="%s", container="%s",container!~"wait-.*", unit="%s"}) by (pod, container, namespace)`,
		p.Namespace, p.Pod, p.Container, p.Unit)
}

// CPU and Momory resources limit (OK)
func BuildResourceLimitQuery(p *PrometheusQueryParams) string {
	return fmt.Sprintf(`avg(kube_pod_container_resource_limits{namespace="%s", pod="%s", container="%s",container!~"wait-.*", unit="%s"}) by (pod, container, namespace)`,
		p.Namespace, p.Pod, p.Container, p.Unit)
}

// Pods running
func BuildRunningPodsQuery(p *PrometheusQueryParams) string {
	var ctnrQuery string
	ctnr := p.Container
	if strings.TrimSpace(ctnr) != "" {
		ctnrQuery = fmt.Sprintf(`, container="%s"`, p.Container)
	}
	return fmt.Sprintf(`kube_pod_container_status_running{instance="%s", namespace="%s"%s}`,
		p.Instance, p.Namespace, ctnrQuery)
}

// Pods by Phase
func BuildPodsByStatusQuery(p *PrometheusQueryParams) string {
	var phaseQuery string
	ctnr := p.Phase
	if strings.TrimSpace(ctnr) != "" {
		phaseQuery = fmt.Sprintf(`, phase="%s"`, p.Phase)
	}
	return fmt.Sprintf(`kube_pod_status_phase{instance="%s", namespace="%s"%s}`,
		p.Instance, p.Namespace, phaseQuery)
}
