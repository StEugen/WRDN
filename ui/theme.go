package ui

import "github.com/rivo/tview"

func SetTableTheme(table *tview.Table) {
    table.SetCell(0, 0, tview.NewTableCell("Node").SetTextColor(tview.ColorYellow).SetAttributes(tview.AttrBold))
    table.SetCell(0, 1, tview.NewTableCell("CPU Usage").SetTextColor(tview.ColorYellow).SetAttributes(tview.AttrBold))
    table.SetCell(0, 2, tview.NewTableCell("Memory Usage").SetTextColor(tview.ColorYellow).SetAttributes(tview.AttrBold))
}

