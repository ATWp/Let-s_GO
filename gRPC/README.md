# [gRPC](https://grpc.io/)
Создадим тестовый gRPC сервер, принимающий 2 числа, складывать и возвращать в качестве ответа их сумму.<br>
gRPC - это framework разработанный в Google для внутренниъ нужд. Cлужит для упрощения написания сервисов|процессов, которые взаимодействую|общаются друг с другом через RPC(Remote Procedure Call - Удалённый вызов процедур).<br>
gRPC работает на http2, описание для методов и прочего Protocol Buffers as JSON(swagger в помощь)<br>
Самое главное - это мультиплатформенность. Файлы методов и настроек для клиента не зависят от языка программирования.<br>


## Client
Клиент будет реализован с помощью библиотеки [Evans](https://github.com/ktr0731/evans)<br>
Evans - gRPC client, которому передается protofile с описаниями сервиса и доступных методов. Он автоматически подключается к серверу и позволяет вызывать описанные методы.

### Install
``` sh
$ go get github.com/ktr0731/evans
```
``` sh
$ go get google.golang.org/protobuf/cmd/protoc-gen-go
```
``` sh
$ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### Protocol Buffers под Windows 10
Скачайте протокол [buffers](https://developers.google.com/protocol-buffers)<br>
1. [Перейти сюда](https://github.com/protocolbuffers/protobuf/releases/)
2. Выберите архив под <br>
    + Windows 10 64 разрядную: protoc-3.13.0-win64.zip<br>
    + Windows 10 32 разрядную: protoc-3.13.0-win32.zip<br>
3. Скачайте, разархивируйте в корневую папку<br>
Например, диск C:\<br>
4. Добавьте путь к файлу protoc.exe в переменную среду PATH 
    + Зайти в проводник
    + Кликнуть на мой компьютер правой кнопкой мыши
    + Свойства
    + Слева: Дополнительные параметры системы
    + В новом окне: Переменные среды
    + В нижнем списке найти Path: нажмите кнопку изменить
    + Справа вверху кнопку: Создать
    + Укажите путь к файлу protoc.exe. У меня он выглядит так: C:\protoc-3.13.0-win64\bin
    + Три раза нажмите ок, Закройте, Перезагрузите ПК<br>

?Не сработало??? Google в помощь :)


### Для компиляции в этом проекте
protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/adder.proto