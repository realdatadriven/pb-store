package routes

import (
	"fmt"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

func getUserToken(e *core.RequestEvent) (string, error) {
	cookie, err := e.Request.Cookie("session")
	if err != nil {
		return "", fmt.Errorf("err getting the session: %s", err)
	}
	return cookie.Value, nil
}

func SSRAuthorizationFromSession(app *pocketbase.PocketBase, se *core.ServeEvent) *router.RouterGroup[*core.RequestEvent] {
	return se.Router.BindFunc(func(e *core.RequestEvent) error {
		token, err := getUserToken(e)
		if err != nil {
			fmt.Println("global middleware check session", err)
		}
		_, err = app.FindAuthRecordByToken(token)
		if err != nil {
			fmt.Println("global middleware check session", err)
		} else {
			//e.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			e.Request.Header.Set("Authorization", fmt.Sprintf("%s", token))
			//fmt.Println(usr, token)
		}
		return e.Next()
	})
}
