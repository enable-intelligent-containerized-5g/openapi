package models

type MlModelTrainingModelInfo struct {
	Name      string  `json:"name"`
	URI       string  `json:"uri"`
	Size      int64   `json:"size"`
	FigureURI string  `json:"figureUri"`
	MSE       float64 `json:"mse"`
	R2        float64 `json:"r2"`
	MSECPU    float64 `json:"mseCpu"`
	R2CPU     float64 `json:"r2Cpu"`
	MSEMem    float64 `json:"mseMem"`
	R2Mem     float64 `json:"r2Mem"`
	MSEThrpt  float64 `json:"mseThrpt"`
	R2Thrpt   float64 `json:"r2Thrpt"`
}
