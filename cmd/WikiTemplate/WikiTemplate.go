// Copyright 2016 Andreas Pannewitz. All rights reserved.

package main

import (
	WT "github.com/GoLangsam/internal/cmd/WikiTemplate/WikiTemplate"
)

func main() {

	WT.Init()

	WT.Wikis.Generate(WT.Wikis)
	WT.Emoticons.Generate(WT.Emoticons)
	WT.Regions.Generate(WT.Regions)
	WT.Countries.Generate(WT.Countries)
	WT.Cities.Generate(WT.Cities)
	WT.Airports.Generate(WT.Cities) // currently via myCities

	if false {
		//	WT.Println(WT.Wikis)
	}
}
