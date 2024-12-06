package models

type PrometheusResult struct {
	Timestamp  float64    `json:"timestamp"`
	Value      float64    `json:"value"`
	MetricType MetricType `json:"metric"`
	Namespace  string     `json:"namespace"`
	Pod        string     `json:"pod"`
	Container  string     `json:"container"`
	Phase      string     `json:"phase"`
	Uid        string     `json:"uid"`
}

func NewPrometheusResult() PrometheusResult {
	return PrometheusResult{
		Timestamp:  0,
		Value:      0,
		MetricType: "",
		Namespace:  "",
		Pod:        "",
		Container:  "",
		Phase:      "",
		Uid:        "",
	}
}