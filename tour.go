package main

import (
	"bufio"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
	"strings"
)

func main() {
	files := []string{
		"welcome.article",
		"basics.article",
		"flowcontrol.article",
		"moretypes.article",
		"methods.article",
		"generics.article",
		"concurrency.article",
	}

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetAliasNbPages("{nb}")
	m.SetPageMargins(20, 10, 20)

	m.RegisterFooter(func() {
		m.Row(10, func() {
			m.Col(12, func() {
				m.Text(fmt.Sprintf("Page %d of {nb}", m.GetCurrentPage()+1), props.Text{
					Size:  10,
					Align: consts.Right,
				})
			})
		})
	})

	for _, fil := range files {
		f, err := os.OpenFile("_content/"+fil, os.O_RDONLY, 444)
		if err != nil {
			panic("Couldn't open file " + fil + ": " + err.Error())
		}

		scan := bufio.NewScanner(f)

		for scan.Scan() {
			text := scan.Text()

			style := consts.Normal
			if strings.HasPrefix(text, "*") {
				style = consts.Bold
			}

			if !strings.HasPrefix(text, ".") &&
				!strings.HasPrefix(text, "#") {
				m.Row(12, func() {
					m.Col(12, func() {
						m.Text(text, props.Text{
							Size:            13,
							Align:           consts.Left,
							Family:          consts.Arial,
							VerticalPadding: 1.2,
							Style:           style,
						})
					})
				})
			}
		}

		if err = scan.Err(); err != nil {
			panic("Error during reading scanner: " + err.Error())
		}

		err = f.Close()
		if err != nil {
			return
		}
		m.AddPage()
	}

	err := os.Mkdir("_pdf", 0755)
	if err != nil && !os.IsExist(err) {
		panic("Couldn't create directory _pdf: " + err.Error())
	}

	err = m.OutputFileAndClose("_pdf/tour.pdf")
	if err != nil {
		panic("Couldn't write tour.pdf: " + err.Error())
	}
}
