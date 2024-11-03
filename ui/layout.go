package ui

import (
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2" 
)

func NewTableLayout() *tview.Table {
    
    table := tview.NewTable().
        SetBorders(true)

    table.SetCell(0, 0, tview.NewTableCell("Node").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
    table.SetCell(0, 1, tview.NewTableCell("CPU Usage").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
    table.SetCell(0, 2, tview.NewTableCell("Memory Usage").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))

    return table
}


