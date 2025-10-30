module github.com/ionos-cloud/sdk-go/v6

go 1.24

require (
	golang.org/x/oauth2 v0.30.0
)

retract (
	v6.0.5851 // Published accidentally.
)
