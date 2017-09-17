// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

import (
	"text/template"
)

// myTemplate holds the external templates
var myTemplate *template.Template

// PopulateTemplate returns the external templates
func PopulateTemplate(pattern string) *template.Template {
	return template.Must(template.ParseGlob(pattern))
}
