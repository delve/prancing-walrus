package middlewares

import (
	"walrusbot/bot/middleware/beforeMiddleware"
)

var Middlewares = []interface{}{
	new(beforeMiddleware.CommandLogging),
	new(beforeMiddleware.CommandAccessControl),
}
