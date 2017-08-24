// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

// Table is the common interface
type Table interface {
	basename() string
	make(len int) Table
	elem(line []string) TableRow
	add(e TableRow)
	//	Generate(myTable Table)
}

type wikis []Wiki
type emoticons []Emoji
type regions []Region
type countries []Country
type currencies []Currency
type cities []CityOf
type airports []Airport

func (t wikis) basename() string      { return "Wiki" }
func (t emoticons) basename() string  { return "Emoji" }
func (t regions) basename() string    { return "GeoRegion" }
func (t countries) basename() string  { return "GeoCountry" }
func (t currencies) basename() string { return "Currency" }
func (t cities) basename() string     { return "CityOf" }
func (t airports) basename() string   { return "Airport" }

func (t wikis) make(len int) wikis           { return make([]Wiki, 0, len) }
func (t emoticons) make(len int) emoticons   { return make([]Emoji, 0, len) }
func (t regions) make(len int) regions       { return make([]Region, 0, len) }
func (t countries) make(len int) countries   { return make([]Country, 0, len) }
func (t currencies) make(len int) currencies { return make([]Currency, 0, len) }
func (t cities) make(len int) cities         { return make([]CityOf, 0, len) }
func (t airports) make(len int) airports     { return make([]Airport, 0, len) }

func (t wikis) elem(line []string) Wiki          { return newWiki(line) }
func (t emoticons) elem(line []string) Emoji     { return newEmoji(line) }
func (t regions) elem(line []string) Region      { return newRegion(line) }
func (t countries) elem(line []string) Country   { return newCountry(line) }
func (t currencies) elem(line []string) Currency { return newCurrency(line) }
func (t cities) elem(line []string) CityOf       { return newCityOf(line) }
func (t airports) elem(line []string) Airport    { return newAirport(line) }

func (t wikis) add(e Wiki)          { t = append(t, e) }
func (t emoticons) add(e Emoji)     { t = append(t, e) }
func (t regions) add(e Region)      { t = append(t, e) }
func (t countries) add(e Country)   { t = append(t, e) }
func (t currencies) add(e Currency) { t = append(t, e) }
func (t cities) add(e CityOf)       { t = append(t, e) }
func (t airports) add(e Airport)    { t = append(t, e) }

func fillfrom(f Table, s [][]string) {
	f = f.make(len(s))
	for _, line := range s {
		var r = f.elem(line)
		if r != nil {
			f.add(r)
		}
	}
}

// ===========================================================================

func newWikis(lines [][]string) wikis {
	var d wikis
	d = d.make(len(lines))
	for _, line := range lines {
		d = append(d, newWiki(line))
	}
	return d
}

func newEmoticons(lines [][]string) emoticons {
	var d = make([]Emoji, 0, len(lines))
	for _, line := range lines {
		d = append(d, newEmoji(line))
	}
	return d
}

func newRegions(lines [][]string) regions {
	var d = make([]Region, 0, len(lines))
	for _, line := range lines {
		d = append(d, newRegion(line))
	}
	return d
}

func newCountries(lines [][]string) countries {
	var d = make([]Country, 0, len(lines))
	for _, line := range lines {
		d = append(d, newCountry(line))
	}
	return d
}

func newCurrencies(lines [][]string) currencies {
	var d = make([]Currency, 0, len(lines))
	for _, line := range lines {
		d = append(d, newCurrency(line))
	}
	return d
}

func newCities(lines [][]string) cities {
	var d = make([]CityOf, 0, len(lines))
	for _, line := range lines {
		d = append(d, newCityOf(line))
	}
	return d
}

func newAirports(lines [][]string) airports {
	var d = make([]Airport, 0, len(lines))
	for _, line := range lines {
		d = append(d, newAirport(line))
	}
	return d
}

// ===========================================================================
