package Interceptor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

/*
	该目录下为自定义拦截器，用于router的pre handle中，可以做一些统计类的操作
	如使用 prometheus 做指标的统计
*/

var (
	metricsTcpReqTotal *prometheus.CounterVec
)

func init() {
	metricsTcpReqTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "tcp_req_total",
		Help: "request total summary by method",
	}, []string{"methodName"})
}

func TcpRequestInterceptor(args ...string) {
	metricsTcpReqTotal.WithLabelValues(args...).Inc()
}
