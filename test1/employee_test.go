package test1

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
	"time"
)

// Позитивный тест
func TestExampleFunctionSuccess(t *testing.T) {
	start := time.Now()

	// Arrange
	a, b := 2, 3
	expected := 5

	// Act
	result := ExampleFunction(a, b)

	// Assert
	duration := time.Since(start)
	writeTestResult("TestExampleFunctionSuccess", fmt.Sprintf("a=%d, b=%d, expected=%d", a, b, expected), result, expected, "", duration)

	if result != expected {
		t.Errorf("Ожидалось: %d, Получено: %d", expected, result)
	}
}

// Негативный тест
func TestExampleFunctionFailure(t *testing.T) {
	start := time.Now()

	// Arrange
	a, b := 2, 2
	unexpected := 5

	// Act
	result := ExampleFunction(a, b)

	// Assert
	duration := time.Since(start)
	var errorMsg string
	if result == unexpected {
		errorMsg = fmt.Sprintf("Результат не должен быть равен %d", unexpected)
	}
	writeTestResult("TestExampleFunctionFailure", fmt.Sprintf("a=%d, b=%d, unexpected=%d", a, b, unexpected), result, unexpected, errorMsg, duration)

	if result == unexpected {
		t.Errorf(errorMsg)
	}
}

// writeTestResult записывает результат теста в CSV-файл
func writeTestResult(testName, input string, actual, expected interface{}, errorMsg string, duration time.Duration) {
	file, err := os.OpenFile("test_results.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Записываем заголовок, если файл пустой
	if fileStat, _ := file.Stat(); fileStat.Size() == 0 {
		writer.Write([]string{"Тест", "Входные данные", "Ожидаемый результат", "Фактический результат", "Ошибка", "Время выполнения"})
	}

	// Записываем результат теста
	writer.Write([]string{
		testName,
		input,
		fmt.Sprintf("%v", expected),
		fmt.Sprintf("%v", actual),
		errorMsg,
		duration.String(),
	})
}
