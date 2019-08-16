# API

## Collections

### GET `/collections`
Получить список всех коллекций.

#### Response
`200 OK`

```json
{
    "items": [<collection>,]
    "total": 3
}
```

### GET `/collection/{collection_id}`
Получить коллекцию по id.

#### Response
`200 OK`

```
<collection>
```

### GET `/collections/random`
Получить случайную коллекцию.

#### Response
`200 OK`

```
<collection>
```

### POST `/collection`
Создать новую коллекцию.

#### Request
```json
{
    "name": "uniq_collection_name"
}
```

#### Response
`201 Created`
```
<collection>
```

### PUT `/collection/{collection_id}`
Обновить коллекцию.

#### Request
```json
{
    "name": "uniq_collection_name"
}
```

#### Response
`200 OK`
```
<collection>
```

### DELETE `/collection/{collection_id}`
Удалить существующую коллекцию.

#### Response
`204 No Content`

### GET `/collection/{collection_id}/words`
Получить список слов из коллекции.

#### Response
`200 OK`

```json
{
    "items": [<word>,]
    "total": 3
}
```

### GET `/word/{word_id}`
Получить слово из коллекции по id.

#### Response
`200 OK`

```
<word>
```

### GET `/collection/{collection_id}/words/random`
Получить случайное слово из коллекции.

#### Response
`200 OK`

```
<word>
```

### GET `/words/random`
Получить случайное слово из коллекции.

#### Response
`200 OK`

```
<word>
```

### POST `/word`
Добавить новое слово в коллекцию.

#### Request
```json
{
    "word": "uniq_word_in_english",
    "translation": "translation_in_russian"
}
```

#### Response
`201 Created`

```
<word>
```

### PUT `/word/{word_id}`
Обновить слово из коллекции по id.

#### Request
```json
{
    "word": "uniq_word_in_english",
    "translation": "translation_in_russian"
}
```

#### Response
`200 OK`

```
<word>
```

### DELETE `/word/{word_id}`
Удалить слово из коллекции.

#### Response
`204 No Content`

## Words

### GET `/words`
Получить список слов из всех слов.

#### Response
`200 OK`

```json
{
    "items": [<word>,]
    "total": 3
}
```

### GET `/word/{word_id}`
Получить слово по id.

#### Response
`200 OK`

```
<word>
```

### GET `/word/random`
Получить случайное слово.

#### Response
`200 OK`

```
<word>
```

### PUT `/word/{word_id}`
Обновить слово по id.

#### Request
```json
{
    "word": "uniq_word_in_english",
    "translation": "translation_in_russian"
}
```

#### Response
`200 OK`

```
<word>
```

### DELETE `/word/{word_id}`
Удалить слово.

#### Response
`204 No Content`