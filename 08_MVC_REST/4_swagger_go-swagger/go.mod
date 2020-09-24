module github.com/spsa/swagger-automation

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/spsa/swagger-automation/api v0.0.0
	github.com/spsa/swagger-automation/docs v0.0.0
)

replace github.com/spsa/swagger-automation/api => ./api

replace github.com/spsa/swagger-automation/docs => ./docs
