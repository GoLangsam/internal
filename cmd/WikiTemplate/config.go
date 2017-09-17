// Copyright 2016 Andreas Pannewitz. All rights reserved.

package WikiTemplate

var myRoot = fullname(myHomePath, "_Common")

const myDataRaw = `data.raw`
const myDataNew = `data.new`

const myExtCSV = `.csv`
const myExtTpl = `.wiki` // `.tmpl`
const myExtWiki = `.wiki`

const myGlobCSV = `*` + myExtCSV
const myGlobTpl = `*` + myExtTpl

var myPathRaw = fullname(myRoot, myDataRaw)
var myPathNew = fullname(myRoot, myDataNew)

//  myTemplatesGlob tells where the external templates are
var myTemplatesGlob = fullname(myPathRaw, myGlobTpl)

// Utf8BOM is written to start of files by Make
var Utf8BOM = []byte{0xEf, 0xBB, 0xBF}

// for csv.ReadAll
const csvComma = '\t'
const csvComment = '#'
const csvLazyQuotes = true
