package utils

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func SendMsgToSentry(c *gin.Context, msg string) {
	if hub := sentrygin.GetHubFromContext(c); hub != nil {
		fmt.Println("sdfsfsdfsdfsdfsfsdfs")
		hub.WithScope(func(scope *sentry.Scope) {
			scope.SetExtra("unwantedQuery", "someQueryDataMaybe")
			// This message will be sent to sentry
			hub.CaptureMessage("Error occured -> " + msg)
		})
	}
}
