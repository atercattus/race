# Гофер гонки

![](techtrain.png)

Нужно написать самое быстрое решение, которое поможет гоферу пройти весь маршрут максимально быстро. Для этого нужно правильно распарсить 
список маршрутов и найти кратчайший путь. Список участков маршрутов в таком формате:

```
A - 3 - B
A - 2 - C
A - 4 - E
B - 8 - C
C - 1 - D
E - 3 - B
```
Каждый участок состоит из двух точек и длинны дороги их связывающей. Имена точек могут быть произвольной длинны, но не могут содержать ничего кроме символов английского алфавита (в том числе заглавных) и цифр. В свою очередь длинна может быть произвольной, но всегда является положительным целым числом. Разделяют их всегда два пробела и тире как на примере выше.

## Задача

Вам нужно построить максимально короткий маршрут от старта до финиша за максимально короткое время. Для этого ваше приложение должно принимать на вход три входных данных - файл "карты", точку старта и точку финиша. После загрузки карты необходимо его разобрать, проверить на корректность (мы ведь не хотим чтобы гофер улетел в пустоту, ведь правда?) и построить тот самый короткий маршрут.

## Куда и как отправлять приложения?

!!!TODO!!!

## Как происходит оценка?

**Магия.** Каждое приложение запускается на нашем сервере с прогревом, после которого приложение несколько раз тестируется на входных данных. По итогам тестов, если отсутствуют сильные отклонения, берется среднее значение которое является итоговым.

## FAQ

### Читать файл, серьезно?
Для такой задачи довольно сложно сделать действительно "чистое" приложение не уйдя в наносекунды(где самые точные процессоры уже начинают терять точность определения) или не получая бинарники размером сотни мегабайт. Мы постарались минимизировать воздействие ввода\вывода на общую картину сделав ее как можно более незаметной. Впрочем, если у вас рекомендации по вопросам "как правильно" - мы с удовольствием вам выслушаем. Только помните, что вариант который вам пришел в голову возможно тоже имеет недостатки. Так или иначе - это тоже считать одним из пунктов по "поиску маршрута" как можно быстро. В конце концов - правильно разворачивать карту нужно тоже уметь :)

### А что если у меня будет набор готовых ответов?
Предсказатель маршрута это хорошо, но насколько это надежно? А вдруг решение которое вы дадите на основе предсказаний будет ошибочным? Гофер отправится в пустоту и ему будет там очень грустно.

А если серьезно - то все "топовые" решения мы конечно проверим на отсутствие "читов". Помните - ваша задача помочь Гоферу найти короткий маршрут, а не решить вопрос любой ценой.

### Как проверяется честность замеров?
Время от времени мы проводим дополнительные тесты.

### А можно характеристики платформы где ищется маршрут?
В вашем распоряжении 4 ядра и 8 гигабайт памяти.

### Я вас поломаю!!!111разраз
Не надо. Пожалуйста.

### А если меня есть еще вопрос?
Подходите к островку GolangRU(22) где мы с радостью вам на него ответим.

До встречи!
