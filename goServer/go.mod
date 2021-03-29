module github.com/heroku/go-getting-started

go 1.12

replace engine => /Engine

require (
	engine v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v0.0.0-20150626140855-4cc2de6207f4 // indirect
	github.com/heroku/x v0.0.0-20171004170240-705849e307dd // indirect
	github.com/labstack/echo/v4 v4.1.17
	github.com/manucorporat/sse v0.0.0-20150604091100-c142f0f1baea // indirect
	gopkg.in/bluesuncorp/validator.v5 v5.9.1 // indirect
)
