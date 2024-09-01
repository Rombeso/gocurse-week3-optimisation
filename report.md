![img.png](img.png)
go tool pprof hw3.test mem.out
(pprof) alloc_space
(pprof) top
![img_1.png](img_1.png)
(pprof) list FastSearch
Основные проблемы
```
4.97MB     28:	fileContents, err := ioutil.ReadAll(file)

1.64MB     1.64MB     38:	lines := strings.Split(string(fileContents), "\n")

512.02kB   512.02kB     42:		user := make(map[string]interface{})

512.28kB        4MB     44:		err := json.Unmarshal([]byte(line), &user)

.    16.54MB     68:			if ok, err := regexp.MatchString("Android", browser); ok && err == nil {

.     9.50MB     90:			if ok, err := regexp.MatchString("MSIE", browser); ok && err == nil {

1MB        1MB    112:		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user["name"], email)

.   514.38kB    115:	fmt.Fprintln(out, "found users:\n"+foundUsers)
```
Заменил на:
28, 38 - читаю в буфер по частям и сохранение построчно в массив
42 - создаю маппу определенной длинны
68, 69 - использую strings.Contains вместо regexp
112 - пишу в буфер
44 - Использую easyjson, добавил структуру с типами, убрал мапу и интерфейсы
Объединил в один цикл поиск разных подстрок
Убрал лишние сохранения данных в новые переменные
![img_2.png](img_2.png)
Убрал сохранение в буффер хранения всех найденных пользователей, сразу выводить результаты
![img_4.png](img_4.png)
![img_5.png](img_5.png)