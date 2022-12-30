package monitor_entry

var (
	LabelNode     = "node"
	LabelCluster  = "cluster"
	LabelApi      = "api"
	LabelApp      = "app"
	LabelHandler  = "handler"
	LabelUpstream = "upstream"
)

var labels = map[string]string{
	LabelApi:      "api_id",
	LabelApp:      "application",
	LabelHandler:  "handler",
	LabelUpstream: "service_id",
}