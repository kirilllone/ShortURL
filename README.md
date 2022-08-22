# ShortURL

2 HandlerFunc: /post, /  

ссылки хранятся в глобальной Map


/post : - запись

полная ссылка передается в теле запроса ("Content-Type: text/plain")

проверяется на соответствие 

генерируется короткая 

200 - если сократилась

500 - если ошибка


короткая ссылка пишется в Response


/ : по короткой возвращает длинную 

полная  сылка передается в теле запроса ("Content-Type: text/plain")

проверяет наличие 

редирект 307

404 - если нет

500 - если ошибка 


полная ссылка пишется в Response



Пример использования 


// curl -v -X POST -H "Content-Type: text/plain" -d "https://yandex.ru" 'localhost:8085/post' (/post)

// curl -v -X POST -H "Content-Type: text/plain" -d "helloworld" 'localhost:8085/post' (неправильный ввод ссылки)

// curl -v -H "Content-Type: text/plain" -d "XVlBz" 'localhost:8085' (поиск по короткой)

// curl -v -H "Content-Type: text/plain" -d "wwwww" 'localhost:8085' (несуществующая короткая)




Скорее всего редирект реализован неправильно, и я не очень понимаю как его правильно отлавливать (статус)

