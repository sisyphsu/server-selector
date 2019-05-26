package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	table1 := widgets.NewTable()
	table1.Rows = [][]string{
		{"header1", "header2", "header3"},
		{"你好吗", "Go-lang is so cool", "Im working on Ruby"},
		{"2016", "10", "11"},
	}
	table1.TextStyle = ui.NewStyle(ui.ColorWhite)
	table1.SetRect(0, 0, 60, 10)

	ui.Render(table1)

	table2 := widgets.NewTable()
	table2.Rows = [][]string{
		{"header1", "header2", "header3"},
		{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		{"2016", "11", "11"},
	}
	table2.TextStyle = ui.NewStyle(ui.ColorWhite)
	table2.TextAlignment = ui.AlignCenter
	table2.RowSeparator = false
	table2.SetRect(0, 10, 20, 20)

	ui.Render(table2)

	table3 := widgets.NewTable()
	table3.Rows = [][]string{
		{"header1", "header2", "header3"},
		{"AAA", "BBB", "CCC"},
		{"DDD", "EEE", "FFF"},
		{"GGG", "HHH", "III"},
	}
	table3.TextStyle = ui.NewStyle(ui.ColorWhite)
	table3.RowSeparator = true
	table3.BorderStyle = ui.NewStyle(ui.ColorGreen)
	table3.SetRect(0, 30, 70, 20)
	table3.FillRow = true
	table3.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	table3.RowStyles[2] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
	table3.RowStyles[3] = ui.NewStyle(ui.ColorYellow)

	ui.Render(table3)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		default:
			println(e.Type, e.ID)
		}
	}
}
