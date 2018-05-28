package main

import (
	"github.com/biz/mdc"
	"github.com/biz/mdc/form"
	"github.com/biz/templates"
	"github.com/gin-gonic/gin"
)

func init() {
	templates.AddPartial("base", `
<!DOCTYPE html>
<html>
	<head>
		<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
		<link rel="stylesheet" href="/static/css/mdc.css" />
		<style>
		/* Setup them colors */
		:root {
			--mdc-theme-primary: #9c27b0;
			--mdc-theme-secondary: #ffab40;
			--mdc-theme-background: #fff;
		}
		</style>
	</head>
	<body>
		{{ block "body" . }}
		{{ end }}
		<script src="/static/js/mdc.js"></script>
		<script>mdc.autoInit();</script>
	</body>
</html>
	`)

	templates.AddView("button", `
{{ define "body" }}
	{{ .mdc.Button "Submit" (.mdc.AddClass "foo" "bar") (.mdc.AddID "foo-bar") (.mdc.AddAttr "foo" "bar") }}
	{{ .mdc.Button "Dense" .mdc.DenseButton }}
	{{ .mdc.Button "Compact" .mdc.CompactButton }}
	{{ .mdc.Button "Submit" .mdc.UnelevatedButton }}
	{{ .mdc.Button "Submit" .mdc.NotRaisedButton }}
	{{ .mdc.IconButton "mood" }}
	{{ .mdc.IconButton "mood" (.mdc.AddTextAfter "After!") }}
	{{ .mdc.IconButton "mood" (.mdc.AddText "Before!") }}
	<br>
	<br>
	{{ .mdc.Fab "add" }}
	{{ .mdc.FabMini "edit" }}

	<section>
		{{ .mdc.TextInput "fname" "First Name" "fname" .mdc.Required }}
	</section>
	<section>
		{{ .mdc.TextInput "disabled" "Disabled" "disabled" .mdc.FieldDisabled }}
	</section>
	<section>
	{{ .mdc.TextInput "lname" "Last Name" "lname" (.mdc.FieldHelperText "Min Length 8") }}
	</section>
	<section>
	{{ .mdc.TextInput "street" "Street" "street" .mdc.FieldBox }}
	</section>
	<section>
	   	{{ .mdc.TextInput "city" "City" "city" .mdc.Required .mdc.FieldOutlined .mdc.FieldDense (.mdc.FieldHelperText "Foo bar baz!") }}
	</section>
	<section>
	   	{{ .mdc.Textarea "desc" "Description" "desc" }}
	</section>
	<section>
	   	{{ .mdc.TextInput "fwidth" "Full Width Input" "fwidth" .mdc.Fullwidth }}
	   	{{ .mdc.Textarea "fwidtharea" "Full Width Textarea" "fwidtharea" .mdc.Fullwidth (.mdc.FieldHelperText "Full width helper")}}
	</section>

	<section>
	   	{{ .mdc.TextInput "leading-icon-outline" "Leading Icon" "leading-icon-outline" .mdc.FieldOutlined (.mdc.FieldLeadingIcon "event") (.mdc.FieldHelperText "Foo bar baz!") }}
	</section>

	<section>
	   	{{ .mdc.TextInput "leading-icon-box" "Leading Icon" "leading-icon-box" .mdc.FieldBox (.mdc.FieldLeadingIcon "event") (.mdc.FieldHelperText "Foo bar baz!") }}
	</section>

	<section>
	   	{{ .mdc.TextInput "trailing-icon-outline" "Trailing Icon" "trailing-icon-outline" .mdc.FieldOutlined (.mdc.FieldTrailingIcon "event") (.mdc.FieldHelperText "Foo bar baz!") }}
	</section>

	<section>
	   	{{ .mdc.TextInput "trailing-icon-box" "Trailing Icon" "trailing-icon-box" .mdc.FieldBox (.mdc.FieldTrailingIcon "event") (.mdc.FieldHelperText "Foo bar baz!") }}
	</section>

	<section>
		{{ .mdc.Checkbox "checkbox" "Checkbox" "checkbox" }}
	</section>

	<section>
		{{ .mdc.Checkbox "checked-checkbox" "Checked Checkbox" "checked-checkbox" }}
	</section>

	<section>
		{{ .mdc.Checkbox "disabled-checkbox" "Disabled Checkbox" "disabled-checkbox" .mdc.CheckboxDisabled }}
	</section>
	
	<fieldset>
		<h5>Climbing Styles</h5>

		{{ .mdc.Checkbox "climbing" "Sport" "sport" (.mdc.InputValue "sport") }}
		{{ .mdc.Checkbox "climbing" "Trad" "trad" (.mdc.InputValue "trad") }}
		{{ .mdc.Checkbox "climbing" "Bouldering" "bouldering" (.mdc.InputValue "bouldering") }}
	</fieldset>

	<fieldset>
		<h5>What is the most fun?</h5>
		{{ .mdc.Radio "most-fun" "Soccer" "soccer-id" "soccer" }}
		{{ .mdc.Radio "most-fun" "Baseball" "baseball-id" "baseball" }}
		{{ .mdc.Radio "most-fun" "Bouldering" "bouldering-id" "bouldering" }}
	</fieldset>
{{ end }}
`)
	/*
	   	<br>
	   	<br>
	   {{ end }}
	   	`)
	*/
}

func main() {
	templates.Parse()

	e := gin.New()

	e.Static("/static/", "./static/")

	e.Use(func(ctx *gin.Context) {
		p := form.New()
		p.SetFormValue("lname", "Maloney")
		p.SetFormValue("fwidtharea", "Some Awesome Stuff!")
		p.SetFormError("fname", "Required Field")
		p.SetFormError("city", "City Please")
		p.SetFormValue("checked-checkbox", "on")

		p.SetFormValue("climbing", "sport")
		p.SetFormValue("climbing", "bouldering")

		p.SetFormValue("most-fun", "bouldering")

		mdc := mdc.New(p)
		ctx.Set("mdc", mdc)
	})

	e.GET("/", func(ctx *gin.Context) {
		templates.MustExecute(ctx.Writer, "base", "button", ctx.Keys)
	})

	e.Run(":8888")
}
