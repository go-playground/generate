Package generate
================

[![GoDoc](https://godoc.org/github.com/go-playground/generate?status.svg)](https://godoc.org/github.com/go-playground/generate)

Package generate runs go generate recursively on a specified path or environment 
variable like $GOPATH and can filter by regex

#### Why would I use it?

When ready to compile your application, you may need to do a bunch of setup, 
run some scripts or even embed static resources and instead of programming
that all into a build script just add all that into //go:generate statements
then run this to recursively go through and run all the setup for you.

#### Why was it created?

To be the best friend of [https://github.com/go-playground/statics](https://github.com/go-playground/statics) which 
automatically embeds the go:generate statement that embeds the static 
resources, but it could be used for anything.

Installation
------------
Use go get.

	go get github.com/go-playground/generate

or to update

	go get -u github.com/go-playground/generate

Then import the validator package into your own code.

	import "github.com/go-playground/generate"

Usage and documentation
------

Please see https://godoc.org/github.com/go-playground/generate for detailed usage docs.

#### Example Usage
* NOTE: this would be from a bash prompt, hence the escaping

generate -i=$GOPATH -ignore=/\\. -match=/github.com/MyOrganizationOrUser 

run generate -h for all options

License
------
Distributed under MIT License, please see license file in code for more details.