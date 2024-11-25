package domain

type LoadTester interface {
	RunLoadTest(config TestConfig) TestResult
}
