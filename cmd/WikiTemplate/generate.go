// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

import (
	"fmt"
)

// Generate generates
func (a wikis) Generate(myWiki []Wiki) {
	nam := asWiki(a.basename())
	out := Make(nam)
	defer out.Close()

	tmpl := a.basename() + "List"
	fmt.Println(tmpl)

	checkerr(myTemplate.ExecuteTemplate(out, tmpl, myWiki), nam)

	tmpl = a.basename() + "Page"
	fmt.Println(tmpl)

	for _, e := range myWiki {
		nam := asWiki(a.basename() + e.ID)
		out := Make(nam)
		defer out.Close()
		checkerr(myTemplate.ExecuteTemplate(out, tmpl, e), nam)
	}
}

// Generate generates
func (a emoticons) Generate(myEmoji []Emoji) {
	nam := asWiki(a.basename())
	out := Make(nam)
	defer out.Close()

	tmpl := a.basename() + "List"
	fmt.Println(tmpl)

	checkerr(myTemplate.ExecuteTemplate(out, tmpl, myEmoji), nam)

	tmpl = a.basename() + "Page"
	fmt.Println(tmpl)

	for _, e := range myEmoji {
		nam := asWiki(a.basename() + e.ID)
		out := Make(nam)
		defer out.Close()
		checkerr(myTemplate.ExecuteTemplate(out, tmpl, e), nam)
	}
}

// Generate generates
func (a cities) Generate(myCities []CityOf) {
	nam := asWiki(a.basename())
	out := Make(nam)
	defer out.Close()

	tmpl := a.basename() + "List"
	fmt.Println(tmpl)

	checkerr(myTemplate.ExecuteTemplate(out, tmpl, myCities), nam)

	tmpl = a.basename() + "Page"
	fmt.Println(tmpl)

	for _, e := range myCities {
		nam := asWiki(a.basename() + e.ID)
		out := Make(nam)
		defer out.Close()
		checkerr(myTemplate.ExecuteTemplate(out, tmpl, e), nam)
	}
}

// Generate generates - Note: Airport currently feeds from myCities!
func (a airports) Generate(myCities []CityOf) {
	nam := asWiki(a.basename())
	out := Make(nam)
	defer out.Close()

	tmpl := a.basename() + "List"
	fmt.Println(tmpl)

	checkerr(myTemplate.ExecuteTemplate(out, tmpl, myCities), nam)

	tmpl = a.basename() + "Page"
	fmt.Println(tmpl)

	for _, e := range myCities {
		if e.Airport.ID != "" {
			nam := asWiki(a.basename() + e.Airport.ID)
			out := Make(nam)
			defer out.Close()
			checkerr(myTemplate.ExecuteTemplate(out, tmpl, e), nam)
		}
	}
}

// Generate generates
func (a countries) Generate(myCountries []Country) {
	nam := asWiki(a.basename())
	out := Make(nam)
	defer out.Close()

	tmpl := a.basename() + "List"
	fmt.Println(tmpl)

	checkerr(myTemplate.ExecuteTemplate(out, tmpl, myCountries), nam)

	tmpl = a.basename() + "Page"
	fmt.Println(tmpl)

	for _, e := range myCountries {
		nam := asWiki(a.basename() + e.ID)
		out := Make(nam)
		defer out.Close()
		checkerr(myTemplate.ExecuteTemplate(out, tmpl, e), nam)
	}
}

// Generate generates
func (a regions) Generate(myRegions []Region) {
	nam := asWiki(a.basename())
	out := Make(nam)
	defer out.Close()

	tmpl := a.basename() + "List"
	fmt.Println(tmpl)

	checkerr(myTemplate.ExecuteTemplate(out, tmpl, myRegions), nam)

	tmpl = a.basename() + "Page"
	fmt.Println(tmpl)

	for _, e := range myRegions {
		nam := asWiki(a.basename() + e.ID)
		out := Make(nam)
		defer out.Close()
		checkerr(myTemplate.ExecuteTemplate(out, tmpl, e), nam)
	}
}
