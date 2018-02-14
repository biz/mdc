package mdc

import (
	"html/template"

	"golang.org/x/net/html"

	"github.com/PuerkitoBio/goquery"
)

func newRadio(name, label, id string) *goquery.Document {
	doc := goquery.NewDocumentFromNode(&html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{{
			Key: "class",
			Val: "mdc-form-field",
		}},
	})
	doc.AppendHtml(`
		<div class="mdc-radio">
			<input class="mdc-radio__native-control" type="radio" id="` + id + `" name="` + name + `">
			<div class="mdc-radio__background">
				<div class="mdc-radio__outer-circle"></div>
				<div class="mdc-radio__inner-circle"></div>
			</div>
		</div>
		<label for="` + id + `">` + label + `</label>
	`)

	return doc
}

func (mdc *MDC) Radio(name, label, id, value string, configs ...ElementConfig) template.HTML {
	doc := newRadio(name, label, id)

	for _, c := range configs {
		c(doc)
	}

	val := ""
	if mdc.Form != nil {
		val = mdc.FormValue(name)
	}

	if val == value {
		doc.Find("input").SetAttr("checked", "checked")
	}

	return render(doc)
}

func (mdc *MDC) RadioDisabled() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.Find("mdc-radio").AddClass("mdc-radio--disabled")
		doc.Find("input").SetAttr("disabled", "disabled")
	}
}
