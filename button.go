package mdc

import (
	"html/template"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func newButton() *goquery.Document {
	return goquery.NewDocumentFromNode(&html.Node{
		Type: html.ElementNode,
		Data: "button",
		Attr: []html.Attribute{{
			Key: "class",
			Val: "mdc-button mdc-js-button mdc-button--raised",
		}},
	})
}

// Button renders a button
func (mdc *MDC) Button(text string, configs ...ElementConfig) template.HTML {
	doc := newButton()
	mdc.AddText(text)(doc)

	for _, c := range configs {
		c(doc)
	}

	return render(doc)
}

// IconButton creates an icon button
func (mdc *MDC) IconButton(icon string, configs ...ElementConfig) template.HTML {
	doc := newButton()
	doc.RemoveClass("mdc-button--raised")

	doc.AppendHtml(`<i class="material-icons mdc-button__icon">` + icon + `</i>`)

	for _, c := range configs {
		c(doc)
	}

	return render(doc)
}

/*
* Button decorate functions
 */

// UnelevatedButton removes the raised button class
func (mdc *MDC) UnelevatedButton() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.RemoveClass("mdc-button---raised")
		doc.AddClass("mdc-button--unelevated")
	}
}

// NotRaisedButton removes the raised button class
func (mdc *MDC) NotRaisedButton() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.RemoveClass("mdc-button--raised")
	}
}

func (mdc *MDC) DenseButton() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.AddClass("mdc-button--dense")
	}
}

// CompactButton removes the raised button class
func (mdc *MDC) CompactButton() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.AddClass("mdc-button--compact")
	}
}
