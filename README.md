
# stwmmensa
Fetch the menu from the *Studentenwerk München* website.

[![Build Status](https://travis-ci.org/folixg/stwmmensa.svg?branch=master)](https://travis-ci.org/folixg/stwmmensa)



### Usage
```
stwmmensa accepts the following arguments:
-h --help: Print this usage info and quit
-o --output: set the path to the output file (mandatory)
-l --location: set mensa location id (default: 421)
-f --format: select format (default: xml)

** Location
Here is a list of valid location codes with the name of the corresponding mensa:
411 : Mensa Leopoldstraße
412 : Mensa Martinsried
421 : Mensa Arcisstraße
422 : Mensa Garching
423 : Mensa Weihenstephan
431 : Mensa Lothstraße
432 : Mensa Pasing

** Format
xml : a generic xml file is created
lis : generate html snippet for LIS-infoscreen

** Examples:
stwmmensa -l 411 -o /my/path/leopold.xml -f xml
stwmmensa --location=423 --output=weihenstephan.html --format=lis
```
### Documentation
The documentation [for the command](https://godoc.org/github.com/folixg/stwmmensa)
and [for the package](https://godoc.org/github.com/folixg/stwmmensa/mensa) are
available on [godoc.org](https://godoc.org)

### LIS specifics
The format `lis` creates a HTML snippet, that is used for the infoscreen in the kitchen at [TUM LIS](http://www.lis.ei.tum.de). *It is not valid HTML for standalone usage.*

In order to cross-compile for the Raspberry Pi run:

```GOOS=linux GOARCH=arm go build -v github.com/folixg/stwmmensa```
