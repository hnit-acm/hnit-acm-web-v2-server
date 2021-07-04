package services

import (
	"github.com/gin-gonic/contrib/sessions"
	"log"
)

var (
	_SessionStore sessions.Store
)

func SessionStore() sessions.Store {
	if _SessionStore == nil {
		s, err := sessions.NewRedisStore(5, "tcp", "127.0.0.1:6379", "", []byte("secret"))
		if err != nil {
			log.Panicln(err)
		}
		_SessionStore = s
	}
	return _SessionStore
}
