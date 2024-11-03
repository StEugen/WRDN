package k8s

import (
    "context"
    "k8s.io/client-go/kubernetes"
    "k8s.io/metrics/pkg/apis/metrics/v1beta1"
    "k8s.io/client-go/rest"
    metricsclientset "k8s.io/metrics/pkg/client/clientset/versioned"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodeMetrics(clientset *kubernetes.Clientset) (*v1beta1.NodeMetricsList, error) {
    
    config, err := rest.InClusterConfig()
    
    if err != nil {
        return nil, err
    }

    metricsClient, err := metricsclientset.NewForConfig(config)
    if err != nil {
        return nil, err
    }

    return metricsClient.MetricsV1beta1().NodeMetricses().List(context.Background(), metav1.ListOptions{})
}
