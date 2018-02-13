package mdc

import (
	"html/template"

	"golang.org/x/net/html"

	"github.com/PuerkitoBio/goquery"
)

func newCheckbox(name, label, id string) *goquery.Document {
	doc := goquery.NewDocumentFromNode(&html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{{
			Key: "class",
			Val: "mdc-form-field",
		}},
	})
	doc.AppendHtml(`
		<div class="mdc-checkbox">
			<input name="` + name + ` "type="checkbox" id="` + id + `" class="mdc-checkbox__native-control"/>
			<div class="mdc-checkbox__background">
				<svg class="mdc-checkbox__checkmark" viewBox="0 0 24 24">
					<path class="mdc-checkbox__checkmark-path" fill="none" stroke="white" d="M1.73,12.91 8.1,19.28 22.79,4.59"></path>
				</svg>
				<div class="mdc-checkbox__mixedmark"></div>
			</div>
		</div>
		<label for="` + id + `">` + label + `</label>
	`)
	return doc
}

func (mdc *MDC) Checkbox(name, label, id string, configs ...ElementConfig) template.HTML {
	doc := newCheckbox(name, label, id)

	for _, c := range configs {
		c(doc)
	}

	value, _ := doc.Find("input").Attr("value")

	values := []string{}
	if mdc.Form != nil {
		values = mdc.FormValues(name)
	}

	if len(value) == 0 {
		// handle checkbox that does not have a value attribute
		if len(values) == 1 && values[0] == "on" {
			doc.Find("input").SetAttr("checked", "checked")
		}
	} else {
		// handle a checkbox that has a value attribute
		for _, val := range values {
			if val == value {
				doc.Find("input").SetAttr("checked", "checked")
				break
			}
		}
	}

	return render(doc)
}

func (mdc *MDC) CheckboxDisabled() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.Find(".mdc-checkbox").AddClass("mdc-checkbox--disabled")
		doc.Find("input").SetAttr("disabled", "disabled")
	}
}

func (mdc *MDC) InputValue(value string) func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.Find("input").SetAttr("value", value)
	}
}
