package ui

import "github.com/rivo/tview"

func SetTableTheme(table *tview.Table) {
    // Use a color for the text
    yellowTextStyle := tview.NewTableCell("Node").SetTextColor(tview.Styles.PrimaryTextColor) // Set color
    table.SetCell(0, 0, yellowTextStyle)

    yellowTextStyle = tview.NewTableCell("CPU Usage").SetTextColor(tview.Styles.PrimaryTextColor) // Set color
    table.SetCell(0, 1, yellowTextStyle)

    yellowTextStyle = tview.NewTableCell("Memory Usage").SetTextColor(tview.Styles.PrimaryTextColor) // Set color
    table.SetCell(0, 2, yellowTextStyle)
}

