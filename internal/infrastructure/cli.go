package infrastructure

import (
	"flag"

	"gitlab.com/douglasschantz/stress-test/internal/domain"
)

type CLI interface {
	Execute()
}
type cli struct {
	loadTester      domain.LoadTester
	reportGenerator domain.ReportGenerator
}

func NewCLI(loadTester domain.LoadTester, reportGenerator domain.ReportGenerator) CLI {
	return &cli{
		loadTester:      loadTester,
		reportGenerator: reportGenerator,
	}
}
func (c *cli) Execute() {
	var config domain.TestConfig
	flag.StringVar(&config.URL, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&config.TotalRequests, "requests", 0, "Número total de requests")
	flag.IntVar(&config.Concurrency, "concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()
	result := c.loadTester.RunLoadTest(config)
	c.reportGenerator.GenerateReport(result)
}
