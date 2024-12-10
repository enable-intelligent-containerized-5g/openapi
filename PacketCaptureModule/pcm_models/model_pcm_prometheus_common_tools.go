package pcm_models

func FindPodByContainer(pods []PrometheusResult, container string) *PrometheusResult {
	for _, pod := range pods {
		if pod.Container == container {
			return &pod // Return Pod
		}
	}
	return nil // Return nil
}