package gui

import (
	"regexp"

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
		"":    {"ec2"},
		"ec2": {"instance", "snapshot(coming soon...)"},
	}
	regionData = []string{"Ohio(us-east-2)", "N. Virginia(us-east-1)", "N. California(us-west-1)",
		"Oregon(us-west-2)", "Cape Town(af-south-1)", "Hong Kong(ap-east-1)",
		"Singapore(ap-southeast-1)", "Sydney(ap-southeast-2)", "Jakarta(ap-southeast-3)",
		"Mumbai(ap-south-1)", "Tokyo(ap-northeast-1)", "Seoul(ap-northeast-2)", "Osaka(ap-northeast-3)",
		"Central(ca-central-1)", "Frankfurt(eu-central-1)"}
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
		target.SetRegion(findRegionId(s))
		d := target.DescribeTarget()
		c.makeTable(d)
		updateContent(c.makeContainer())
	})
	c.regions = regionEntry
}

func (c *ServiceContent) makeContainer() *container.Split {
	split := container.NewHSplit(c.tree, container.NewBorder(c.regions, nil, nil, nil, c.table))
	split.Offset = 0.35
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

	w.Resize(fyne.NewSize(800, 500))
	w.ShowAndRun()
}

func updateContent(c *container.Split) {
	w.SetContent(c)
}

func findRegionId(s string) string {
	regionMathcer := regexp.MustCompile(`.*\((.*)\)`)
	group := regionMathcer.FindSubmatch([]byte(s))
	return string(group[1])
}
