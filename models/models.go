package models

type Code struct {
	Language string `json:"lang"`
	Code     string `json:"code"`
}

type CodeAnalysisResponse struct {
	CodeOptimizationLevel  int     `json:"code_optimization_level"`
	CPUPerformance         string  `json:"cpu_performance"`
	MemoryUsage            string  `json:"memory_usage"`
	Error                  *string `json:"error"` // null if no error
	Output                 string  `json:"output"`
	RedundantBlock         *string `json:"redundant_block"`          // repeated/unnecessary code
	UnusedVariables        *string `json:"unused_variables"`         // null if none
	UnusedFunctions        *string `json:"unused_functions"`         // null if none
	SuggestedOptimizedCode string  `json:"suggested_optimized_code"` // fully optimized code
}
