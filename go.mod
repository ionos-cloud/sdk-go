module github.com/ionos-cloud/sdk-go/v6

go 1.21

require (
	golang.org/x/oauth2 v0.23.0
)

retract (
	v6.0.5851 // Published accidentally.
)
