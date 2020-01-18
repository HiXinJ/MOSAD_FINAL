客户端用到的接口
- [注册](#%e6%b3%a8%e5%86%8c)
  - [body param](#body-param)
  - [Response](#response)
- [登录](#%e7%99%bb%e5%bd%95)
  - [body param](#body-param-1)
  - [Response](#response-1)
- [获取用户新词](#%e8%8e%b7%e5%8f%96%e7%94%a8%e6%88%b7%e6%96%b0%e8%af%8d)
  - [Query Param](#query-param)
  - [Response](#response-2)
- [标记单词已学](#%e6%a0%87%e8%ae%b0%e5%8d%95%e8%af%8d%e5%b7%b2%e5%ad%a6)
  - [Request Header](#request-header)
  - [Response](#response-3)
- [获取复习单词](#%e8%8e%b7%e5%8f%96%e5%a4%8d%e4%b9%a0%e5%8d%95%e8%af%8d)
  - [Query Param](#query-param-1)
  - [Response](#response-4)
- [打卡](#%e6%89%93%e5%8d%a1)
  - [Request Header](#request-header-1)
  - [Response](#response-5)
- [单词翻译](#%e5%8d%95%e8%af%8d%e7%bf%bb%e8%af%91)
  - [Query Param](#query-param-2)
  - [Response](#response-6)
- [上传用户头像](#%e4%b8%8a%e4%bc%a0%e7%94%a8%e6%88%b7%e5%a4%b4%e5%83%8f)
  - [Request Header](#request-header-2)
  - [Response](#response-7)
- [获取用户头像](#%e8%8e%b7%e5%8f%96%e7%94%a8%e6%88%b7%e5%a4%b4%e5%83%8f)
  - [Request Header](#request-header-3)
  - [Response](#response-8)


Response至少包含"error_message"和"message".    
如果 error_message == "" 代表一切正常,否则有错误。


## 注册
POST /users

### body param
```json
{
 	"user_name":"string",
	"password":"string"
}
```

### Response

```json
{
   "error_message": "string",
   "message": "string",
   "token":"string",
   "user_id":"int",
}
```


## 登录
POST /users/login

### body param
```json
{
 	"user_name":"string",
	"password":"string",
    "user_id":"int",
}
```

### Response

```json
{
   "error_message": "string",
   "message": "string",
   "token":"string"
}
```


## 获取用户新词
GET /words/new?user_name=&size=

### Query Param
- user_name 用户名，必须
- size 单词个数，可选 默认5

### Response
```json
{
    "error_message": "",
    "fanyi": [
        {
            "query": "division",
            "us_phonetic": "dɪˈvɪʒn",
            "uk_phonetic": "dɪˈvɪʒn",
            "explains": [
                "n. [数] 除法；部门；分配；分割；师（军队）；赛区"
            ],
            "examples_ch": "…德国在二战结束时分裂成两个国家，之后又统一了。",
            "examples_en": "...the unification of Germany, after its division into two states at the end of World War Two. "
        },
        {
            "query": "flicker",
            "us_phonetic": "ˈflɪkər",
            "uk_phonetic": "ˈflɪkə(r)",
            "explains": [
                "n. (Flicker) （美、澳、英）弗利克（人名）",
                "v. 闪烁，摇曳；颤动；（情绪等）闪现；快速瞥视；扑动翅膀",
                "n. 闪烁，闪光；霎时的感情（犹豫、激动等）；抖动，颤动；（电影的）图像闪烁；微小动作；（鸟）扑翅鴷属"
            ],
            "examples_ch": "荧光灯闪了闪，接着房间里就亮得令人目眩了。",
            "examples_en": "Fluorescent lights flickered, and then the room was blindingly bright. "
        },
    ],
    "message": "success"
}
```

## 标记单词已学
POST /words/learnedword?word=  

需token

### Request Header 
```json
{
    "authorization": "string",
}
```


### Response
```json
{
    "error_message": "",
    "message": "success"
}
```

## 获取复习单词
GET /words/reviews?size=

需token

### Query Param
- size 单词个数，可选 默认5

### Response

例如，`GET http://localhost:8081/words/reviews?size=1`

```json
{
    "error_message": "",
    "fanyi": [
        {
            "query": "flicker",
            "us_phonetic": "ˈflɪkər",
            "uk_phonetic": "ˈflɪkə(r)",
            "explains": [
                "n. (Flicker) （美、澳、英）弗利克（人名）",
                "v. 闪烁，摇曳；颤动；（情绪等）闪现；快速瞥视；扑动翅膀",
                "n. 闪烁，闪光；霎时的感情（犹豫、激动等）；抖动，颤动；（电影的）图像闪烁；微小动作；（鸟）扑翅鴷属"
            ],
            "examples_ch": "荧光灯闪了闪，接着房间里就亮得令人目眩了。",
            "examples_en": "Fluorescent lights flickered, and then the room was blindingly bright. "
        },
    ],
    "message": "success"
}
```

## 打卡
POST /users/{user_name}/daka  


需要token
### Request Header 
```json
{
    "authorization": "string",
}
```

### Response  
Response包含所有打过卡的日期date，以及总共打卡天数ndays.
```json
{
    "date": [
        {
            "Year": 2020,
            "Month": 1,
            "Day": 16
        }
    ],
    "error_message": "",
    "message": "success",
    "ndays": 1
}
```

## 单词翻译

GET /words/translation?word=

### Query Param
- word 要查询的单词

### Response
例如 `GET /words/translation?word=tile`
```json
{
    "error_message": "",
    "message": "success",
    "translation": [
        {
            "query": "tile",
            "us_phonetic": "taɪl",
            "uk_phonetic": "taɪl",
            "explains": [
                "n. 瓷砖，瓦片",
                "vt. 铺以瓦；铺以瓷砖",
                "n. (Tile)人名；(俄、塞、萨摩)蒂勒"
            ],
            "examples_ch": "埃米走过走廊时，鞋子在地砖上吱吱作响。",
            "examples_en": "Amy's shoes squeaked on the tiles as she walked down the corridor. "
        }
    ]
}
```

## 上传用户头像

POST /users/head

### Request Header 
```json
{
    "authorization": "string",
}
```

### Response
```json
{
    "error_message": "",
    "message": "success"
}
```

## 获取用户头像
GET /users/head

### Request Header 
```json
{
    "authorization": "string",
}
```

### Response
图像字节流
