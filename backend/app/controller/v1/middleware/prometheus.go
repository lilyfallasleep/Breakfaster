package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"breakfaster/pkg/stat"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = ""

// EndpointList is a list of endpoints
var EndpointList = []string{
	"/metrics",
	"/callback",
	"/api/v1/foods",
	"/api/v1/employee/line-uid",
	"/api/v1/employee/emp-id",
	"/api/v1/employee",
	"/api/v1/order",
	"/api/v1/order/pick",
	"/api/v1/orders",
	"/api/v1/next-week",
}
var (
	labels = []string{"status", "endpoint", "method"}

	uptime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "uptime",
			Help:      "HTTP service uptime.",
		}, nil,
	)

	reqCountPerEndpoint = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "http_request_count_total_per_endpoint",
			Help:      "Total number of HTTP requests made per endpoint.",
		}, labels,
	)

	reqCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "http_request_count_total",
			Help:      "Total number of HTTP requests made.",
		}, nil,
	)

	userCPU = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "user_cpu_usage",
			Help:      "User CPU Usage.",
		}, nil,
	)

	systemCPU = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "system_cpu_usage",
			Help:      "System CPU Usage.",
		}, nil,
	)

	memUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "system_mem_usage",
			Help:      "System Mem Usage.",
		}, nil,
	)

	diskUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "system_disk_usage",
			Help:      "System Disk Usage.",
		}, nil,
	)
)

// init registers the prometheus metrics
func init() {
	prometheus.MustRegister(uptime, reqCount, reqCountPerEndpoint, userCPU, systemCPU, memUsage, diskUsage)
	go recordServerMetrics()
}

func recordServerMetrics() {
	prev := stat.GetServer()
	for range time.Tick(time.Second) {
		cur := stat.GetServer()
		cpuTotal := float64(cur.CPU.Total - prev.CPU.Total)
		uptime.WithLabelValues().Inc()
		userCPU.WithLabelValues().Set(float64(cur.CPU.User-prev.CPU.User) / cpuTotal * 100)
		systemCPU.WithLabelValues().Set(float64(cur.CPU.System-prev.CPU.System) / cpuTotal * 100)
		memUsage.WithLabelValues().Set(cur.Mem.Usage() * 100)
		diskUsage.WithLabelValues().Set(cur.Disk.Usage() * 100)
		prev = cur
	}
}

// calcRequestSize returns the size of request object
func calcRequestSize(r *http.Request) float64 {
	size := 0
	if r.URL != nil {
		size = len(r.URL.String())
	}

	size += len(r.Method)
	size += len(r.Proto)

	for name, values := range r.Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}
	size += len(r.Host)

	// r.Form and r.MultipartForm are assumed to be included in r.URL.
	if r.ContentLength != -1 {
		size += int(r.ContentLength)
	}
	return float64(size)
}

// PromMiddleware returns the prometheus middleware handler
func PromMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		status := fmt.Sprintf("%d", c.Writer.Status())
		endpoint := c.Request.URL.Path
		method := c.Request.Method
		recordEndpoint := ""
		for i := range EndpointList {
			if strings.HasPrefix(endpoint, EndpointList[i]) && len(recordEndpoint) < len(EndpointList[i]) {
				recordEndpoint = EndpointList[i]
			}
		}
		if recordEndpoint == "" {
			recordEndpoint = "unknown"
		}
		lvs := []string{status, recordEndpoint, method}

		reqCount.WithLabelValues().Inc()
		reqCountPerEndpoint.WithLabelValues(lvs...).Inc()
	}
}

// PromHandler return http handler for prometheus
func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
