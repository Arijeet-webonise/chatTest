package app

import (
	"github.com/Arijeet-webonise/chatTest/pkg/logger"
	"github.com/Arijeet-webonise/chatTest/pkg/session"
	"github.com/Arijeet-webonise/chatTest/pkg/storage"
	"github.com/Arijeet-webonise/chatTest/pkg/templates"
	"github.com/go-zoo/bone"
)

// App container for the application
type App struct {
	Router         *bone.Mux
	TplParser      templates.TplParse
	SessionManager session.SessionManager
	Log            logger.ILogger
	StorageManager storage.StorageManager
}
