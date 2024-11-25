package service

import (
	"fmt"

	"gitlab.com/douglasschantz/stress-test/internal/domain"
)

type reportService struct{}

func NewReportService() domain.ReportGenerator {
	return &reportService{}
}
func (s *reportService) GenerateReport(result domain.TestResult) {
	fmt.Printf("Tempo total gasto: %v\n", result.Duration)
	fmt.Printf("Total de requests: %d\n", result.TotalRequests)
	fmt.Printf("Requests com status 200: %d\n", result.StatusCounts[200])

	for status, count := range result.StatusCounts {
		if status != 200 {
			fmt.Printf("Requests com status %v: %d\n", status, count)
		}
	}
}
