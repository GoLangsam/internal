// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

// TableRow is the common interface
type TableRow interface {
	new()
}

// ===========================================================================
type id struct{ ID string }
type description struct{ Name string }
type alias struct{ Alias string }

// ===========================================================================
func newWiki(line []string) Wiki {
	var e Wiki
	e.ID = line[0]
	e.Icon = line[1]
	return e
}

// Wiki are wikis
type Wiki struct {
	id
	Icon            string
	PageTitlePrefix string
	FavName         string
	FavOpt          string
	FavRelpath      string
}

// ===========================================================================
func newEmoji(line []string) Emoji {
	var e Emoji
	e.ID = unBracket(line[0], "[", "]")
	e.Name = line[1]
	e.Alias = e.Name
	e.Aliases = splitEmoji(line[2], " ")
	//	drop = line[3]
	e.ASCIIes = splitEmoji(unBracket(line[4], "<< ", ">>"), " ")
	return e
}

// Emoji are emoticons
type Emoji struct {
	id
	description
	alias
	Aliases []string
	ASCIIes []string
}

// ===========================================================================
func newCurrency(line []string) Currency {
	var e Currency
	e.ID = line[0]
	e.Name = line[1]
	e.Alias = line[2]
	e.Country.ID = line[3]
	return e
}

// Currency are Currencies of Countries
type Currency struct {
	id
	description
	alias
	Country
}

// ===========================================================================
func newRegion(line []string) Region {
	var e Region
	e.ID = line[0]
	e.Name = line[1]
	e.Alias = line[2]
	return e
}

// Region are Regions of Countries
type Region struct {
	id
	description
	alias
}

// ===========================================================================
func newCountry(line []string) Country {
	var e Country
	e.ID = line[0]
	e.Name = line[1]
	e.Currency.ID = line[2]
	e.Region.ID = line[3]
	e.Alias = line[4]
	return e
}

// Country are Countries
type Country struct {
	id
	description
	Currency id
	Region
	alias
}

// ===========================================================================
func newCityOf(line []string) CityOf {
	var e CityOf
	e.Country.ID = line[0]
	e.Airport.ID = line[1]
	e.ID = line[2]
	e.Name = line[2]
	e.Alias = line[3]
	return e
}

// CityOf are cities
type CityOf struct {
	id
	description
	alias
	Country
	Airport id
}

// ===========================================================================
func newAirport(line []string) Airport {
	var e Airport
	e.ID = line[0]
	e.Name = line[1]
	e.Alias = line[2]
	e.Country.ID = line[3]
	e.CityOf.ID = line[4]
	return e
}

// Airport are airports
type Airport struct {
	id
	description
	alias
	Country
	CityOf
}
