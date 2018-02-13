package mdc

import (
	"html/template"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func newFab(icon string) *goquery.Document {
	return goquery.NewDocumentFromNode(&html.Node{
		Type: html.ElementNode,
		Data: "button",
		Attr: []html.Attribute{{
			Key: "class",
			Val: "mdc-fab material-icons",
		}},
		FirstChild: &html.Node{
			Type: html.ElementNode,
			Data: "span",
			Attr: []html.Attribute{{
				Key: "class",
				Val: "mdc-fab__icon",
			}},
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: icon,
			},
		},
	})
}

func (mdc *MDC) Fab(icon string) template.HTML {
	doc := newFab(icon)
	return render(doc)
}

func (mdc *MDC) FabMini(icon string) template.HTML {
	doc := newFab(icon)
	doc.AddClass("mdc-fab--mini")
	return render(doc)
}
