package domain

type ReportGenerator interface {
	GenerateReport(result TestResult)
}
