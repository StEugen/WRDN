package ui

import (
    "wrdn/k8s"
    "github.com/rivo/tview"
    "k8s.io/client-go/kubernetes"
    "log"
    "time"
)

func UpdateMetrics(table *tview.Table, clientset *kubernetes.Clientset) {
    for {
        time.Sleep(5 * time.Second)
        nodeMetrics, err := k8s.GetNodeMetrics(clientset)
        if err != nil {
            log.Println("Error fetching node metrics:", err)
            continue
        }

        for i, metric := range nodeMetrics.Items {
            table.SetCell(i+1, 0, tview.NewTableCell(metric.Name))
            table.SetCell(i+1, 1, tview.NewTableCell(metric.Usage.Cpu().String()))
            table.SetCell(i+1, 2, tview.NewTableCell(metric.Usage.Memory().String()))
        }
    }
}

