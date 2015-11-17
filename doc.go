/*
Package generate runs go generate recursively on a specified path or environment
variable like $GOPATH and can filter by regex

Why would I use it?

	When ready to compile your application, you may need to do a bunch of setup,
	run some scripts or even embed static resources and instead of programming
	that all into a build script just add all that into //go:generate statements
	then run this to recursively go through and run all the setup for you.

Why was it created?

	To be the best friend of https://github.com/go-playground/statics which
	automatically embeds the go:generate statement that embeds the static
	resources, but it could be used for anything.

Example Usage
	* NOTE: this would be from a bash prompt, hence the escaping

	generate -i=$GOPATH -ignore=/\\. -match=/github.com/MyOrganizationOrUser

	run generate -h for all options

*/
package main
