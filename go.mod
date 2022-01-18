module mingo/audit

go 1.18

// +heroku goVersion 1.17 

require github.com/gorilla/mux v1.8.0

require (
	github.com/mattn/go-sqlite3 v1.14.10
	github.com/stretchr/testify v1.7.0
	gopkg.in/go-playground/validator.v9 v9.31.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.1.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/mingo/swagger/go v1.0.0 => ./go
