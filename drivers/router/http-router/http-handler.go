package http_router

import (
	"net/http"

	http_context2 "github.com/eolinker/apinto/node/http-context"

	http_complete "github.com/eolinker/apinto/drivers/router/http-router/http-complete"
	"github.com/eolinker/apinto/service"

	"github.com/eolinker/eosc/eocontext"
	http_context "github.com/eolinker/eosc/eocontext/http-context"
)

var completeCaller = http_complete.NewHttpCompleteCaller()

type httpHandler struct {
	completeHandler eocontext.CompleteHandler

	routerName  string
	serviceName string

	finisher  eocontext.FinishHandler
	service   service.IService
	filters   eocontext.IChainPro
	disable   bool
	websocket bool
}

func (h *httpHandler) ServeHTTP(ctx eocontext.EoContext) {
	httpContext, err := http_context.Assert(ctx)
	if err != nil {
		return
	}
	if h.disable {
		httpContext.Response().SetStatus(http.StatusNotFound, "")
		httpContext.Response().SetBody([]byte("router disable"))
		httpContext.FastFinish()
		return
	}
	if h.websocket {
		wsCtx, err := http_context2.NewWebsocketContext(httpContext)
		if err != nil {
			httpContext.Response().SetStatus(http.StatusInternalServerError, "")
			httpContext.Response().SetBody([]byte(err.Error()))
			httpContext.FastFinish()
			return
		}
		ctx = wsCtx
	}

	//Set Label
	ctx.SetLabel("api", h.routerName)
	ctx.SetLabel("service", h.serviceName)
	ctx.SetLabel("ip", httpContext.Request().ReadIP())
	ctx.SetFinish(h.finisher)
	ctx.SetCompleteHandler(h.completeHandler)
	ctx.SetApp(h.service)
	ctx.SetBalance(h.service)
	ctx.SetUpstreamHostHandler(h.service)

	h.filters.Chain(ctx, completeCaller)
	ctx.GetFinish().Finish(ctx)
}
