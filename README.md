# Сокращение ссылок
Сервис написан на фремворке fiber, запись в базу PostgreSQL осуществлялась через GORM. В качестве in-memory хранилища используется map. Токен короткой ссылки генерируются библиотекой go-nanoid


## **Запуск in-memory или в БД Postgres зависит от значения поля DB_PG в .env файле**

**true - запуск в Postgres**

**false - запуск in-memory**

**(идет проверка на значение поля, любое отличное от true / false значение зафаталит приложение)**



**PostgreSQL**

Поле DB_PG=true в .env файле

`docker-compose up`

`go run cmd/main.go`

После запуска сервиса в Postgres видим успешные логи о создании конфига из .env файла, успешное подключение к БД
![Screenshot from 2023-05-02 14-30-48](https://user-images.githubusercontent.com/101155101/235654542-2a942fe8-71b5-4ea9-b72a-817b5d334bad.png)

**В Postman выставляем метод POST, ссылка на апи localhost:8080/api/create**
Тело запроса - полный URL. 

![Screenshot from 2023-05-03 01-31-21](https://user-images.githubusercontent.com/101155101/235799639-69bf5638-766d-46ed-bd3b-4816bc624b4f.png)


Ответ - сокращенная ссылка + информация о добавлении в базу + код 200

![Screenshot from 2023-05-03 01-32-32](https://user-images.githubusercontent.com/101155101/235799834-9570d106-1680-4661-8d08-260dab07c173.png)



Точно так же добавляем еще 1 ссылку

![Screenshot from 2023-05-03 01-34-04](https://user-images.githubusercontent.com/101155101/235800052-a9131c6b-a59e-40d8-8017-da48362ed7b6.png)

Ответ - сокращенная ссылка + информация о добавлении в базу + код 200

![Screenshot from 2023-05-03 01-33-55](https://user-images.githubusercontent.com/101155101/235800065-dd2d9a95-8fa5-4152-b7e2-881c9b1e5836.png)


В БД появилась информация об этих 2 запросах (полная ссылка и сокращенная) 

![Screenshot from 2023-05-03 01-36-25](https://user-images.githubusercontent.com/101155101/235800339-912ca147-5710-4a64-930a-4db59ea8e994.png)




**В Postman выставляем метод GET, ссылка на апи:**
`localhost:8080/api/find/` + `полученный токен из базы`
Подставляем нашу короткую ссылку, полученную ранее

![Screenshot from 2023-05-03 01-42-23](https://user-images.githubusercontent.com/101155101/235801106-ea49ba5b-df2c-414c-95a2-64be47808672.png)

Получаем ответ в виде полной ссылки и http код 200

![Screenshot from 2023-05-03 01-42-31](https://user-images.githubusercontent.com/101155101/235801141-61fbb692-9983-4655-9f71-3f98fcf10045.png)


## In-memory

Поле DB_PG=false в .env файле

`docker-compose up`

`go run cmd/main.go`


![Screenshot from 2023-05-02 14-48-10](https://user-images.githubusercontent.com/101155101/235657885-ce65c82d-164e-48e9-acef-ff47a1da0f99.png)

**В Postman выставляем метод POST, ссылка на апи localhost:8080/api/create**
Тело запроса - полный URL. 


![Screenshot from 2023-05-03 01-54-35](https://user-images.githubusercontent.com/101155101/235802745-16604f4b-cd4a-422b-bce7-4a8d758c3ff1.png)



Получаем ответ в виде короткой ссылки и http кода 200

![Screenshot from 2023-05-03 01-54-41](https://user-images.githubusercontent.com/101155101/235802755-0d63fda8-1d8a-4c1e-8aff-c3f4a67ead12.png)



Добавим еще 1 url 

![Screenshot from 2023-05-03 01-55-01](https://user-images.githubusercontent.com/101155101/235802761-690b5ecb-222f-4228-b9e2-6249544304ae.png)



Ответ 

![Screenshot from 2023-05-03 01-55-05](https://user-images.githubusercontent.com/101155101/235802766-1add0551-18eb-4676-b541-df702a03bbb5.png)




**В Postman выставляем метод GET, ссылка на апи:**
`localhost:8080/api/find/` + `полученный токен из базы`
Подставляем нашу короткую ссылку, полученную ранее

![Screenshot from 2023-05-03 01-58-52](https://user-images.githubusercontent.com/101155101/235803163-218d8ae9-a02f-4cba-ab64-2210f41e622b.png)


получаем ответ в виде полной ссылки и http код 200

![Screenshot from 2023-05-03 01-56-39](https://user-images.githubusercontent.com/101155101/235802939-eae9e060-263a-4ab3-afac-8bb727bcb513.png)


При GET запросе выводится в консоль мапа со всеми записями URL в ней (короткий url - ключ, полный url - значение)

![Screenshot from 2023-05-03 01-56-56](https://user-images.githubusercontent.com/101155101/235802963-3f06f440-b1e9-46f7-96d5-37b5eb4513cb.png)


