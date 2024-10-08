package health

import (
	"fmt"

	"github.com/PrathamSkilltelligent/pmgingo/routeutils"
	"github.com/PrathamSkilltelligent/pmtestgo/service/apis/config"
	"github.com/PrathamSkilltelligent/pmtestgo/service/apis/models/health"
	"github.com/samber/mo"
)

func HealthController(appContext *config.AppContext, c *routeutils.RequestCtx) mo.Result[*health.HealthResponse] {
	healthRes := health.NewHealthResponse("{{APP_NAME}} is alive.")
	fmt.Println("{{My_App}} is alive!!!")
	return mo.Ok(
		&healthRes,
	)
}
