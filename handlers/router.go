package handlers

import (
	//"github.com/hulucat/utils"
	"net/http"
)

type HttpRouter struct {
}

func (router HttpRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//实际业务逻辑
	switch r.URL.Path {
	case "/init":
		handleInit(w, r)
	case "/status":
		handleGetStatus(w, r)
	case "/build":
		handleBuild(w, r)
	case "/destroy":
		handleDestroy(w, r)
	case "/result":
		handleGetResult(w, r)
	case "/oauth_redirect":
		handleOauthRedirect(w, r)
	}
}
