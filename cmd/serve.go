package cmd

import (
	"io"

	"clevergo.tech/authmiddleware"
	"clevergo.tech/clevergo"
	"clevergo.tech/form"
	"clevergo.tech/jetpackr"
	"clevergo.tech/jetrenderer"
	"clevergo.tech/log"
	"clevergo.tech/osenv"
	"github.com/CloudyKit/jet/v4"
	"github.com/alexedwards/scs/v2"
	"github.com/gobuffalo/packr/v2"
	"pkg.razonyang.com/go-auth0-web-app/internal/core"
	"pkg.razonyang.com/go-auth0-web-app/internal/handlers/dashboard"
	"pkg.razonyang.com/go-auth0-web-app/internal/handlers/home"
	"pkg.razonyang.com/go-auth0-web-app/internal/handlers/user"
	"pkg.razonyang.com/go-auth0-web-app/internal/middleware"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func init() {
	app.Commands = append(app.Commands, serveCmd)
}

var serveCmd = &cli.Command{
	Name:  "serve",
	Usage: "start a HTTP server",
	Action: func(c *cli.Context) error {
		clevergo.SetLogger(provideLogger())
		app := clevergo.New()
		app.Decoder = form.New()
		sessionManager := provideSessionManager()
		app.Use(
			core.ErrorHandler,
			clevergo.WrapHH(sessionManager.LoadAndSave),
			authmiddleware.New(core.NewSessionAuthenticator(sessionManager)),
			middleware.IsAuthenticated("/login", clevergo.PathSkipper("/", "/login", "/callback", "/assets/*")),
		)
		app.Renderer = provideRenderer()

		app.ServeFiles("/assets/*filepath", packr.New("public", "../public"))

		handlers := []core.Handler{
			home.New(),
			user.New(sessionManager),
			dashboard.New(),
		}
		for _, handler := range handlers {
			handler.Register(app)
		}

		return app.Run(osenv.Get("HTTP_ADDR", ":8080"))
	},
}

func provideRenderer() clevergo.Renderer {
	box := packr.New("views", "../views")
	set := jet.NewHTMLSetLoader(jetpackr.New(box))
	set.SetDevelopmentMode(core.IsDevelopMode())
	renderer := jetrenderer.New(set)
	renderer.SetBeforeRender(func(w io.Writer, name string, vars jet.VarMap, data interface{}, c *clevergo.Context) error {
		ctx := c.Context()
		vars.Set("user", authmiddleware.GetIdentity(ctx))
		return nil
	})
	return renderer
}

func provideSessionManager() *scs.SessionManager {
	m := scs.New()
	m.Cookie.HttpOnly = !core.IsDevelopMode()
	return m
}

func provideLogger() log.Logger {
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}
