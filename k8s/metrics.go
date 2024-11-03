package k8s

import (
    "context"
    "k8s.io/client-go/kubernetes"
    "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

func GetNodeMetrics(clientset *kubernetes.Clientset) (*v1beta1.NodeMetricsList, error) {
    metricsClient := clientset.MetricsV1beta1()
    return metricsClient.NodeMetricses().List(context.Background(), v1.ListOptions{})
}

