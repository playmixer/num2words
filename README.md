## num2words

num2words - Конвертация чисел в слова на Go (Golang)

## Как использовать 
```go
import github.com/playmixer/num2words
```
Ковертируем числа
```go
    str := num2words.Convert(1002) // output "одна тысяча два"
    ...
    str := num2words.Convert(-5) // output "минус пять"
    ...
    str := num2words.Convert(320) // output "триста двадцать"

```