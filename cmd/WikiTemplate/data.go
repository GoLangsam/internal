// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

// MyRawData holds the raw data
var MyRawData [][]string

// Wikis are all of Wiki
var Wikis wikis

// Emoticons are all of Emoji
var Emoticons emoticons

// Regions are all of Regions
var Regions regions

// Countries are all Country
var Countries countries

// Currencies are all of Currency
var Currencies currencies

// Cities are all CityOf
var Cities cities

// Airports are all Airport
var Airports airports

//Init data
func Init() {
	// myTemplate holds the external templates
	myTemplate = PopulateTemplate(myTemplatesGlob)

	Wikis = newWikis(read(asCSV(Wikis.basename())))
	//	Wikis = fillfrom(Wikis, read(asCSV(Wikis.basename())))
	Emoticons = newEmoticons(read(asCSV(Emoticons.basename())))
	Regions = newRegions(read(asCSV(Regions.basename())))
	Countries = newCountries(read(asCSV(Countries.basename())))
	Cities = newCities(read(asCSV(Cities.basename())))
	if false {
		Currencies = newCurrencies(read(asCSV(Currencies.basename())))
		Airports = newAirports(read(asCSV(Airports.basename())))
	}
}
