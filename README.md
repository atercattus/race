# Гофер гонки

Нужно накодить самую быструю машину, которая пройдет весь маршрут максимально быстро. Для этого нужно правильно распарсить 
список маршрутов и найти крадчайший путь. Список маршрутов в таком формате:

```
A - 3 - B
A - 2 - C
A - 4 - E
B - 8 - C
C - 1 - D
A - 4 - E
E - 3 - D
```

Вам нужно построить максимально короткий маршрут от точки `A` до точки `D` за максимально короткое время

Весь ваш код должен быть в файле `faster_car.go`

Для запуска тестов используем команду:

```
go test -bench=. -v
```

