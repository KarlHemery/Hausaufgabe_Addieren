package services

type DBInterface interface {
	CalculateSum(num1, num2 float64) (float64, error)
	Close()
}
