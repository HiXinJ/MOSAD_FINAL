

## 下载不安装

```bash
go get -d github.com/hixinj/MOSAD_FINAL
cd $GOPATH/src/github.com/hixinj/MOSAD_FINAL
go run main.go
```

## 接口
GET /words?size=2

默认size=5

### 返回样例
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


