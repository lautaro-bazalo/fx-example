package server

import (
	"context"
	"fmt"
	"fxdemo/internal/pkg/user"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"net/http"
)

func ServerRun(gEngine *gin.Engine, lc fx.Lifecycle) {

	port := viper.GetString("app.port")

	server := http.Server{
		Addr:    port,
		Handler: gEngine,
	}

	lc.Append(
		fx.Hook{OnStart: func(ctx context.Context) error {
			fmt.Print("Running on port", port)
			go func() {
				if err := server.ListenAndServe(); err != nil {
					fmt.Errorf("failed to lsiten and serve from server: %v", err)
				}
			}()
			return nil
		},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		},
	)
}

func RegisterService(gEngine *gin.Engine, router user.Router) {
	gRoute := gEngine.Group("")
	router.Handler(gRoute)
}
