package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atsushi-kitazawa/aws-service-viewer/aws"
)

type ServiceContent struct {
	tree    *fyne.Container
	regions *widget.Select
	table   *widget.Table
}

var (
	w        fyne.Window
	target   = aws.NewTarget()
	treeData = map[string][]string{
		"": {"EC2"},
	}
	regionData = []string{"us-east-2", "us-east-1", "us-west-1", "us-west-2", "af-south-1", "ap-east-1",
		"ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-south-1",
		"ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ca-central-1", "eu-central-1"}
)

func (c *ServiceContent) makeTree() {
	tree := widget.NewTreeWithStrings(treeData)
	tree.OnSelected = func(uid string) {
		target.SetService(uid)
	}
	c.tree = container.NewBorder(nil, nil, nil, nil, tree)
}

func (c *ServiceContent) makeTable(data [][]string) {
	table := widget.NewTable(
		func() (int, int) {
			return len(data), 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})
	table.SetColumnWidth(0, 150)
	table.SetColumnWidth(1, 150)
	table.SetColumnWidth(2, 200)
	c.table = table
}

func (c *ServiceContent) makeRegionsEntry() {
	regionEntry := widget.NewSelect(regionData, func(s string) {
		target.SetRegion(s)
		d := target.DescribeTarget()
		c.makeTable(d)
		updateContent(c.makeContainer())
	})
	c.regions = regionEntry
}

func (c *ServiceContent) makeContainer() *container.Split {
	split := container.NewHSplit(c.tree, container.NewBorder(c.regions, nil, nil, nil, c.table))
	split.Offset = 0.2
	return split
}

func Run() {
	a := app.New()
	w = a.NewWindow("aws service viewer")

	content := &ServiceContent{}
	content.makeTree()
	content.makeTable([][]string{})
	content.makeRegionsEntry()

	updateContent(content.makeContainer())

	w.Resize(fyne.NewSize(750, 500))
	w.ShowAndRun()
}

func updateContent(c *container.Split) {
	w.SetContent(c)
}
