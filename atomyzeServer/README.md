
<span  style="font-size: medium; color: aqua " >

## Тестовое задание Go Atomyze
#### Цель

Разработать сервис для создания и проверки подписи аргументов.

Сервис должен быть описать в формате grpc/protobuf.
Должна использоваться криптография ECDSA.
Нужно написать тесты на методы сервиса, с успешными и неуспешными сценариями

#### Методы сервиса
В сервисе должно быть два метода:
## <span  style="font-size: big; color: aquamarine" >CreateTransaction</span>

### Входные параметры:

args [ ][ ]byte
private key [ ]byte

#### В методе:

нужно создать хэш (любой хэш функцией) от  args (входной параметр)
подписать хэш ( создать подпись ) приватным ключем, переданным во входных параметрах
создать публичный ключ на базе приватного

#### Выходные параметры:

hash [ ]byte
public key  [ ]byte
signature [ ]byte
## <span  style="font-size: big; color: aquamarine" >VerifyTransaction</span>
#### Входные параметры:

hash [ ]byte
public key  [ ]byte
signature [ ]byte

#### В методе:

нужно проверить что подпись соответствует хэшу и публичному ключу

#### Выходные параметры:

signature valid bool

</span>
