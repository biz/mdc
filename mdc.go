// Package mdc is used to create Material Design Web Components in an html template
package mdc

import (
	"html/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/biz/bufpool"
	"golang.org/x/net/html"
)

var buffers = bufpool.New()

type ElementConfig func(*goquery.Document)

func New(f Form) *MDC {
	mdc := &MDC{}
	mdc.Form = f
	return mdc
}

type Form interface {
	FormValue(string) string
	FormError(string) string
	FormValues(string) []string
}

type MDC struct {
	Form
}

/*
* Helper functions used to decorate
 */

// Ripple adds a ripple effect
func (mdc *MDC) Ripple() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.First().SetAttr("data-mdc-auto-init", "MDCRipple")
	}
}

// AddID adds an html element ID
func (mdc *MDC) AddID(id string) func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.First().SetAttr("id", id)
	}
}

// AddClass adds html element classes
func (mdc *MDC) AddClass(classes ...string) func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.First().AddClass(classes...)
	}
}

// AddAttr adds html attribute
func (mdc *MDC) AddAttr(key, val string) func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.SetAttr(key, val)
	}
}

// AddText adds text to the element
func (mdc *MDC) AddText(text string) func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.PrependNodes(&html.Node{
			Type: html.TextNode,
			Data: text,
		})
	}
}

// AddTextAfter adds text to the element
func (mdc *MDC) AddTextAfter(text string) func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.AppendNodes(&html.Node{
			Type: html.TextNode,
			Data: text,
		})
	}
}

func render(doc *goquery.Document) template.HTML {
	buf := buffers.Get()
	defer buffers.Put(buf)
	for i, n := range doc.Nodes {
		if err := html.Render(buf, n); err != nil {
			panic(err)
		}
		if i > 0 {
			buf.Write([]byte(`
`))
		}
	}
	buf.Write([]byte(`
`))

	return template.HTML(buf.String())
}
