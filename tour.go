package main

import (
	"bufio"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
	"strings"
)

func main() {
	dir, err := os.ReadDir("_content")
	if err != nil {
		panic("Couldn't open directory _content: " + err.Error())
	}
	for _, fil := range dir {
		f, err := os.OpenFile("_content/"+fil.Name(), os.O_RDONLY, 444)
		if err != nil {
			panic("Couldn't open file " + fil.Name() + ": " + err.Error())
		}

		m := pdf.NewMaroto(consts.Portrait, consts.A4)

		m.SetPageMargins(20, 10, 20)

		scan := bufio.NewScanner(f)

		for scan.Scan() {
			text := scan.Text()

			style := consts.Normal

			if strings.HasPrefix(text, "*") {
				style = consts.Bold
			}

			if !strings.HasPrefix(text, ".play") {
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

		_ = os.Mkdir("_pdf", os.ModeDir)

		name := fil.Name()

		if dot := strings.LastIndex(name, "."); dot != -1 {
			name = name[:dot] + ".pdf"
		} else {
			name = name + ".pdf"
		}

		err = m.OutputFileAndClose("_pdf/" + name)
		if err != nil {
			panic("Couldn't write pdf " + name + ":" + err.Error())
		}
	}
}
