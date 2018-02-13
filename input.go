package mdc

import (
	"html/template"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func newTextInput(name, label, id, value string) *goquery.Document {
	doc := goquery.NewDocumentFromNode(&html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{{
			Key: "class",
			Val: "mdc-text-field",
		}, {
			Key: "data-mdc-auto-init",
			Val: "MDCTextField",
		}},
	})

	doc.AppendNodes(
		&html.Node{
			Type: html.ElementNode,
			Data: "input",
			Attr: []html.Attribute{{
				Key: "id",
				Val: id,
			}, {
				Key: "value",
				Val: value,
			}, {
				Key: "name",
				Val: name,
			}, {
				Key: "type",
				Val: "text",
			}, {
				Key: "class",
				Val: "mdc-text-field__input",
			}},
		},
		&html.Node{
			Type: html.ElementNode,
			Data: "label",
			Attr: []html.Attribute{{
				Key: "class",
				Val: "mdc-text-field__label",
			}, {
				Key: "for",
				Val: id,
			}},
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: label,
			},
		},
		&html.Node{
			Type: html.ElementNode,
			Data: "div",
			Attr: []html.Attribute{{
				Key: "class",
				Val: "mdc-line-ripple",
			}},
		},
	)

	return doc
}

func newTextarea(name, label, id, value string) *goquery.Document {
	doc := newTextInput(name, label, id, value)
	doc.AddClass("mdc-text-field--textarea")
	doc.Find("input").Nodes[0].Data = "textarea"
	ta := doc.Find("textarea")
	ta.Nodes[0].FirstChild = &html.Node{
		Type: html.TextNode,
		Data: value,
	}

	// set some defaults
	ta.SetAttr("rows", "4")
	ta.SetAttr("cols", "20")

	// Remove ripple
	doc.Find(".mdc-line-ripple").Remove()
	return doc
}

func (mdc *MDC) TextInput(name, label, id string, configs ...ElementConfig) template.HTML {
	value := ""
	errorMsg := ""
	if mdc.Form != nil {
		value = mdc.FormValue(name)
		errorMsg = mdc.FormError(name)
	}

	doc := newTextInput(name, label, id, value)

	// proper classes to render when
	// value is present
	if len(value) > 0 {
		doc.AddClass("mdc-text-field--upgraded")
		doc.Find("label").AddClass("mdc-text-field__label--float-above")
	}

	// run configs
	for _, c := range configs {
		c(doc)
	}

	// Add field validation errors if present - this will take the place of the helper text if present
	if len(errorMsg) > 0 {
		doc.First().AddClass("mdc-text-field--invalid")
		if len(doc.Nodes) == 1 {
			mdc.FieldHelperText(errorMsg)(doc)
		} else {
			doc.Closest(".mdc-text-field-helper-text").Empty()
			doc.Closest(".mdc-text-field-helper-text").First().Nodes[0].AppendChild(&html.Node{
				Type: html.TextNode,
				Data: errorMsg,
			})
		}
		doc.Last().RemoveClass("mdc-text-field-helper-text--persistent")
		doc.Last().AddClass("mdc-text-field-helper-text--validation-msg")
	}

	return render(doc)
}

func (mdc *MDC) Textarea(name, label, id string, configs ...ElementConfig) template.HTML {
	value := ""
	errorMsg := ""
	if mdc.Form != nil {
		value = mdc.FormValue(name)
		errorMsg = mdc.FormError(name)
	}

	doc := newTextarea(name, label, id, value)

	for _, c := range configs {
		c(doc)
	}

	// Add field validation errors if present - this will take the place of the helper text if present
	if len(errorMsg) > 0 {
		doc.First().AddClass("mdc-text-field--invalid")
		if len(doc.Nodes) == 1 {
			mdc.FieldHelperText(errorMsg)(doc)
		} else {
			doc.Closest(".mdc-text-field-helper-text").Empty()
			doc.Closest(".mdc-text-field-helper-text").First().Nodes[0].AppendChild(&html.Node{
				Type: html.TextNode,
				Data: errorMsg,
			})
		}
		doc.Last().RemoveClass("mdc-text-field-helper-text--persistent")
		doc.Last().AddClass("mdc-text-field-helper-text--validation-msg")
	}

	return render(doc)
}

// FieldHelperText is used to add helper text to an input field
// link: https://github.com/material-components/material-components-web/tree/master/packages/mdc-textfield/helper-text
func (mdc *MDC) FieldHelperText(text string) func(*goquery.Document) {
	return func(doc *goquery.Document) {
		input := doc.Find(".mdc-text-field__input")
		heID, _ := input.Attr("id")
		heID += "-helper-text"
		input.SetAttr("aria-controls", heID)

		// create helper text
		ht := goquery.NewDocumentFromNode(&html.Node{
			Type: html.ElementNode,
			Data: "p",
			Attr: []html.Attribute{{
				Key: "class",
				Val: "mdc-text-field-helper-text mdc-text-field-helper-text--persistent",
			}, {
				Key: "aria-hiddin",
				Val: "true",
			}, {
				Key: "id",
				Val: heID,
			}},
		})
		ht.AppendNodes(&html.Node{
			Type: html.TextNode,
			Data: text,
		})

		doc.Selection = doc.AddNodes(ht.Nodes[0])
	}
}

// Add html attributes to the input field
// NOTE: only works with Elements that have input fields
func (mdc *MDC) InputAttr(key, val string) func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.Find("input").SetAttr(key, val)
	}
}

func (mdc *MDC) Required() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.Find("input").SetAttr("required", "required")
	}
}

func (mdc *MDC) FieldBox() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.AddClass("mdc-text-field--box")
	}
}

func (mdc *MDC) FieldOutlined() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.AddClass("mdc-text-field--outlined")
		doc.Find(".mdc-line-ripple").Remove()
		doc.First().AppendHtml(`
<div class="mdc-text-field__outline">
	<svg>
		<path class="mdc-text-field__outline-path"/>
	</svg>
</div>
<div class="mdc-text-field__idle-outline"></div>
		`)
	}
}

func (mdc *MDC) FieldDense() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.AddClass("mdc-text-field--dense")
	}
}

func (mdc *MDC) FieldDisabled() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.AddClass("mdc-text-field--disabled")
	}
}

func (mdc *MDC) Fullwidth() func(*goquery.Document) {
	return func(doc *goquery.Document) {
		doc.AddClass("mdc-text-field--fullwidth")
		if doc.HasClass("mdc-text-field--textarea") {
			doc.Find("textarea").RemoveAttr("cols")
			return
		}
		label := doc.Find("label")
		input := doc.Find("input")
		input.SetAttr("placeholder", label.Text())
		input.SetAttr("aria-label", label.Text())
		label.Remove()
	}
}

// FieldLeadingIcon add a leading icon
// NOTE: FieldBox or FieldOutlined is required to precede the icon
func (mdc *MDC) FieldLeadingIcon(icon string) func(*goquery.Document) {
	return func(doc *goquery.Document) {
		if !doc.HasClass("mdc-text-field--box") && !doc.HasClass("mdc-text-field--outlined") {
			return
		}

		doc.AddClass("mdc-text-field--with-leading-icon")
		doc.PrependHtml(`<i class="material-icons mdc-text-field__icon" tabindex="0">` + icon + `</i>`)

		if doc.HasClass("mdc-text-field--box") {
			doc.Find(".mcd-line-ripple").Remove()
			doc.AppendHtml(`<div class="mdc-text-field__bottom-line"></div>`)
		}
	}
}

// FieldTrailingIcon add a leading icon
// NOTE: FieldBox or FieldOutlined is required to precede the icon
func (mdc *MDC) FieldTrailingIcon(icon string) func(*goquery.Document) {
	return func(doc *goquery.Document) {
		if !doc.HasClass("mdc-text-field--box") && !doc.HasClass("mdc-text-field--outlined") {
			return
		}

		doc.AddClass("mdc-text-field--with-trailing-icon")
		doc.Find("label").AfterHtml(`<i class="material-icons mdc-text-field__icon" tabindex="0">` + icon + `</i>`)

		if doc.HasClass("mdc-text-field--box") {
			doc.Find(".mcd-line-ripple").Remove()
			doc.AppendHtml(`<div class="mdc-text-field__bottom-line"></div>`)
		}
	}
}
