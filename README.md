# Решение тестового задания на позицию Junior Go Developer

### Основаня задача
Написать функцию, которая первым аргументом принимает массив из структур данных `Transaction` для графика, а вторым необходимый интервал для форматирования. Возвращает массив из структур данных `Transaction`, в котором у каждого элемента значение поля `Value` является самым актуальном во временном отрезке, а поле `Timestamp` округленный до требуемого временного интервала.

### Требования
- Код с решением должен быть загружен на github ✅
- Решение должно быть разработано с помощью Go или Node.js ✅
- Покрытие тестами будет плюсом ✅
- Репозиторий с кодом должен содержать ридми файл, где будет инструкция по запуску приложения ✅

## Инструкция по использованию
Главная функция принимает массив транзакций и интервал согласно заданию, возвращает массив отформатированных транзакций и ошибку.

```go
func FormatTransactionByInterval(transactions []Transaction, interval time.Duration) ([]Transaction, error)
```

Функция для второго аргумента принимает только интервалы равные **месяцу, недели, дню или часу**.
Для второго аргумента предоставлены готовые константные значения для удобного ввода. Доступны следующие константные интервалы:
```go
const (
	Hour  = time.Hour
	Day   = 24 * Hour
	Week  = 7 * Day
	Month = 30 * Day
)
```

Функция проверяет валидность интервала, если переданный интервал не валиден возвращает ошибку `NotValidIntervalError`.

**Пример выполнения**:
```go
txs := []Transaction{
		{Value: 4456, Timestamp: time.Date(2021, 03, 18, 0, 10, 48, 0, time.UTC)},
		{Value: 4231, Timestamp: time.Date(2021, 03, 17, 23, 10, 48, 0, time.UTC)},
		{Value: 5212, Timestamp: time.Date(2021, 03, 17, 22, 10, 48, 0, time.UTC)},
		{Value: 4321, Timestamp: time.Date(2021, 03, 16, 10, 10, 48, 0, time.UTC)},
		{Value: 4567, Timestamp: time.Date(2021, 03, 16, 5, 10, 48, 0, time.UTC)},
	}

formatedTxs, err := FormatTransactionByInterval(txs, Day)
if err != nil {
	fmt.Print(err)
	return
}

for i, tx := range formatedTxs {
	fmt.Printf("%d. Value - %d\n   Date - %v\n", i+1, tx.Value, tx.Timestamp)
}
```
**Результат:**
```bash
1. Value - 4456
   Date - 2021-03-18 00:00:00 +0000 UTC
2. Value - 4231
   Date - 2021-03-17 00:00:00 +0000 UTC
3. Value - 4321
   Date - 2021-03-16 00:00:00 +0000 UTC
```
## Тесты
Функция тестируется на корректное поведение функции при передаче разных интервалах, а также корректное определение не валидных значений переданного интервала.