package main

import (
    "log"
    "wrdn/k8s"
    "wrdn/ui"
    "github.com/rivo/tview"
)

func main() {
    clientset, err := k8s.NewClientset()
    if err != nil {
        log.Fatal(err)
    }

    app := tview.NewApplication()

    table := ui.NewTableLayout()
    
    go ui.UpdateMetrics(table, clientset)

    if err := app.SetRoot(table, true).Run(); err != nil {
        log.Fatal(err)
    }
}


