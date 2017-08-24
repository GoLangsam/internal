// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

func main() {

	Init()

	Wikis.Generate(Wikis)
	Emoticons.Generate(Emoticons)
	Regions.Generate(Regions)
	Countries.Generate(Countries)
	Cities.Generate(Cities)
	Airports.Generate(Cities) // currently via myCities

	if false {
		//	Println(Wikis)
	}
}
