GET /words?size=2

默认size=5


```json
{
    "error_message": "",
    "fanyi": [
        {
            "basic": {
                "explains": [
                    "n. [特医] 解压；降压"
                ],
                "phonetic": "ˌdiːkəmˈpreʃn",
                "uk-phonetic": "ˌdiːkəmˈpreʃn",
                "us-phonetic": "ˌdiːkəmˈpreʃn"
            },
            "errorCode": 0,
            "query": "decompression",
            "web": [
                {
                    "key": "Decompression",
                    "value": [
                        "减压",
                        "解压缩",
                        "减压术"
                    ]
                },
                {
                    "key": "File decompression",
                    "value": [
                        "文件解压缩",
                        "文件解紧缩",
                        "文件"
                    ]
                },
                {
                    "key": "microvascular decompression",
                    "value": [
                        "微血管减压术",
                        "显微血管减压术",
                        "微血管减压"
                    ]
                }
            ]
        },
        {
            "basic": {
                "explains": [
                    "n. 重新油漆的东西；重画的画",
                    "vt. 重画；重新绘制；重漆"
                ],
                "phonetic": "riː'peɪnt",
                "uk-phonetic": "riː'peɪnt",
                "us-phonetic": "'ripent"
            },
            "errorCode": 0,
            "query": "repaint",
            "web": [
                {
                    "key": "repaint",
                    "value": [
                        "重画",
                        "重新髹漆",
                        "重绘"
                    ]
                },
                {
                    "key": "re Repaint",
                    "value": [
                        "重画屏幕"
                    ]
                },
                {
                    "key": "Repaint Publish",
                    "value": [
                        "重绘发帖"
                    ]
                }
            ]
        }
    ],
    "message": "success",
    "words": [
        "decompression",
        "repaint"
    ]
}
```


## 有道API

```bash
curl  http://fanyi.youdao.com/openapi.do?keyfrom=pdblog&key=993123434&type=data&doctype=json&version=1.1&only=dict&q=server
```
```json
{
    "basic": {
        "us-phonetic": "ˈsɜːrvər",
        "phonetic": "ˈsɜːvə(r)",
        "uk-phonetic": "ˈsɜːvə(r)",
        "explains": [
            "n. 发球员；服伺者；服勤者；伺候者；计算机网络服务器；上菜用具；助祭，辅祭；计算器主机；分菜勺",
            "n. (Server) （美、俄、西、法）塞尔韦尔（人名）"
        ]
    },
    "query": "server",
    "errorCode": 0,
    "web": [
        {
            "value": [
                "服务器",
                "发球员",
                "伺服器"
            ],
            "key": "SERVER"
        },
        {
            "value": [
                "文件服务器",
                "档案伺服器",
                "档案服务器"
            ],
            "key": "file server"
        },
        {
            "value": [
                "邮件服务器",
                "邮件伺服器",
                "电子邮件服务器"
            ],
            "key": "mail server"
        }
    ]
}
```