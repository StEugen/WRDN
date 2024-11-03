package ui

import "github.com/rivo/tview"

func NewTableLayout() *tview.Table {
    table := tview.NewTable().
        SetBorders(true).
        SetTitle("Kubernetes Metrics").
        SetTitleAlign(tview.AlignLeft)
    
    table.SetCell(0, 0, tview.NewTableCell("Node"))
    table.SetCell(0, 1, tview.NewTableCell("CPU Usage"))
    table.SetCell(0, 2, tview.NewTableCell("Memory Usage"))
    
    return table
}

