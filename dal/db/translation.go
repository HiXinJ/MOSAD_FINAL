package mydb

type SimpleTranslation struct {
	Query     string   `json:"query"`
	USP       string   `json:"us_phonetic"`
	UKP       string   `json:"uk_phonetic"`
	Explains  []string `json:"explains"`
	ExampleCH string   `json:"examples_ch"`
	ExampleEN string   `json:"examples_en"`
}

type Translation struct {
	Basic     Basic_t `json:"basic"`
	Query     string  `json:"query"`
	ErrorCOde int     `json:"errorCode"`
	Web       []Web_t `json:"web"`
}

type Basic_t struct {
	USP      string   `json:"us-phonetic"`
	P        string   `json:"phonetic"`
	UKP      string   `json:"uk-phonetic"`
	Explains []string `json:"explains"`
}

type Web_t struct {
	Value []string `json:"value"`
	Key   string   `json:key`
}

/*

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
*/
