# Service provider system
Итоговая работа по курсу Skillbox "Go-разработчик"


## Сборка
Используя go

    go build cmd/main.go

Используя make

    make build 

## Сторонние Пакеты
    github.com/spf13/viper - для файлов конфигурации
    github.com/go-chi/chi/v5 - маршрутизатор для http

## Конфигурация
Файл конфигурации [config.yaml](configs/config.yaml) в папке [config](configs)

```yaml
# Настройки основной программы
server:
  addr: "localhost"
  port: "8484"

# Настройки симулятора
simulator:
  addr: "localhost"
  port: "8787"

# Настройки файлов с данными
data:
  path: "data/"
```

## Запуск
Симулятор интегрирован в данный проект. На выбор можно запустить основную программу, либо симулятор.
Используется флаг mode

Для корректной работы нужно запустить симулятор и отдельно запустить основную программу

Запуск симулятора

    ./main -mode="simulator"

Запуск основной программы (по умолчанию)

    ./main -mode="server"


Адрес и порт по умолчанию берется из конфигурационного файла.
Можно переназначить адрес или порт при запуске основной программы. Используется флаг addr и флаг port

    ./main -addr="localhost" -port="8282"


## Запрос
Пример запроса на сервер

    http://localhost:8585/

Пример ответа от сервера

```JSON
{
    "status": true,
    "data": {
        "sms": [
            [
                {
                    "country": "Saint Barthélemy",
                    "bandwidth": "96",
                    "response_time": "215",
                    "provider": "Kildy"
                },
                {
                    "country": "New Zealand",
                    "bandwidth": "76",
                    "response_time": "1526",
                    "provider": "Kildy"
                },
                {
                    "country": "Monaco",
                    "bandwidth": "30",
                    "response_time": "1778",
                    "provider": "Kildy"
                },
                {
                    "country": "United States of America",
                    "bandwidth": "3",
                    "response_time": "515",
                    "provider": "Rond"
                },
                {
                    "country": "Bulgaria",
                    "bandwidth": "32",
                    "response_time": "1237",
                    "provider": "Rond"
                },
                {
                    "country": "Canada",
                    "bandwidth": "1",
                    "response_time": "1793",
                    "provider": "Rond"
                },
                {
                    "country": "Russian Federation",
                    "bandwidth": "70",
                    "response_time": "1808",
                    "provider": "Topolo"
                },
                {
                    "country": "United Kingdom of Great Britain and Northern Ireland",
                    "bandwidth": "65",
                    "response_time": "214",
                    "provider": "Topolo"
                },
                {
                    "country": "France",
                    "bandwidth": "29",
                    "response_time": "167",
                    "provider": "Topolo"
                },
                {
                    "country": "Austria",
                    "bandwidth": "69",
                    "response_time": "1298",
                    "provider": "Topolo"
                },
                {
                    "country": "Denmark",
                    "bandwidth": "70",
                    "response_time": "455",
                    "provider": "Topolo"
                },
                {
                    "country": "Spain",
                    "bandwidth": "67",
                    "response_time": "383",
                    "provider": "Topolo"
                },
                {
                    "country": "Switzerland",
                    "bandwidth": "57",
                    "response_time": "1349",
                    "provider": "Topolo"
                },
                {
                    "country": "Peru",
                    "bandwidth": "62",
                    "response_time": "1041",
                    "provider": "Topolo"
                }
            ],
            [
                {
                    "country": "Austria",
                    "bandwidth": "69",
                    "response_time": "1298",
                    "provider": "Topolo"
                },
                {
                    "country": "Bulgaria",
                    "bandwidth": "32",
                    "response_time": "1237",
                    "provider": "Rond"
                },
                {
                    "country": "Canada",
                    "bandwidth": "1",
                    "response_time": "1793",
                    "provider": "Rond"
                },
                {
                    "country": "Denmark",
                    "bandwidth": "70",
                    "response_time": "455",
                    "provider": "Topolo"
                },
                {
                    "country": "France",
                    "bandwidth": "29",
                    "response_time": "167",
                    "provider": "Topolo"
                },
                {
                    "country": "Monaco",
                    "bandwidth": "30",
                    "response_time": "1778",
                    "provider": "Kildy"
                },
                {
                    "country": "New Zealand",
                    "bandwidth": "76",
                    "response_time": "1526",
                    "provider": "Kildy"
                },
                {
                    "country": "Peru",
                    "bandwidth": "62",
                    "response_time": "1041",
                    "provider": "Topolo"
                },
                {
                    "country": "Russian Federation",
                    "bandwidth": "70",
                    "response_time": "1808",
                    "provider": "Topolo"
                },
                {
                    "country": "Saint Barthélemy",
                    "bandwidth": "96",
                    "response_time": "215",
                    "provider": "Kildy"
                },
                {
                    "country": "Spain",
                    "bandwidth": "67",
                    "response_time": "383",
                    "provider": "Topolo"
                },
                {
                    "country": "Switzerland",
                    "bandwidth": "57",
                    "response_time": "1349",
                    "provider": "Topolo"
                },
                {
                    "country": "United Kingdom of Great Britain and Northern Ireland",
                    "bandwidth": "65",
                    "response_time": "214",
                    "provider": "Topolo"
                },
                {
                    "country": "United States of America",
                    "bandwidth": "3",
                    "response_time": "515",
                    "provider": "Rond"
                }
            ]
        ],
        "mms": [
            [
                {
                    "country": "Saint Barthélemy",
                    "provider": "Kildy",
                    "bandwidth": "48",
                    "response_time": "667"
                },
                {
                    "country": "New Zealand",
                    "provider": "Kildy",
                    "bandwidth": "92",
                    "response_time": "1063"
                },
                {
                    "country": "Monaco",
                    "provider": "Kildy",
                    "bandwidth": "66",
                    "response_time": "840"
                },
                {
                    "country": "United States of America",
                    "provider": "Rond",
                    "bandwidth": "50",
                    "response_time": "919"
                },
                {
                    "country": "Bulgaria",
                    "provider": "Rond",
                    "bandwidth": "60",
                    "response_time": "817"
                },
                {
                    "country": "Canada",
                    "provider": "Rond",
                    "bandwidth": "18",
                    "response_time": "531"
                },
                {
                    "country": "Türkiye",
                    "provider": "Rond",
                    "bandwidth": "86",
                    "response_time": "1066"
                },
                {
                    "country": "Russian Federation",
                    "provider": "Topolo",
                    "bandwidth": "82",
                    "response_time": "830"
                },
                {
                    "country": "United Kingdom of Great Britain and Northern Ireland",
                    "provider": "Topolo",
                    "bandwidth": "92",
                    "response_time": "1271"
                },
                {
                    "country": "France",
                    "provider": "Topolo",
                    "bandwidth": "87",
                    "response_time": "66"
                },
                {
                    "country": "Austria",
                    "provider": "Topolo",
                    "bandwidth": "75",
                    "response_time": "631"
                },
                {
                    "country": "Denmark",
                    "provider": "Topolo",
                    "bandwidth": "85",
                    "response_time": "1544"
                },
                {
                    "country": "Spain",
                    "provider": "Topolo",
                    "bandwidth": "55",
                    "response_time": "790"
                },
                {
                    "country": "Switzerland",
                    "provider": "Topolo",
                    "bandwidth": "45",
                    "response_time": "1921"
                },
                {
                    "country": "Peru",
                    "provider": "Topolo",
                    "bandwidth": "2",
                    "response_time": "36"
                }
            ],
            [
                {
                    "country": "Austria",
                    "provider": "Topolo",
                    "bandwidth": "75",
                    "response_time": "631"
                },
                {
                    "country": "Bulgaria",
                    "provider": "Rond",
                    "bandwidth": "60",
                    "response_time": "817"
                },
                {
                    "country": "Canada",
                    "provider": "Rond",
                    "bandwidth": "18",
                    "response_time": "531"
                },
                {
                    "country": "Denmark",
                    "provider": "Topolo",
                    "bandwidth": "85",
                    "response_time": "1544"
                },
                {
                    "country": "France",
                    "provider": "Topolo",
                    "bandwidth": "87",
                    "response_time": "66"
                },
                {
                    "country": "Monaco",
                    "provider": "Kildy",
                    "bandwidth": "66",
                    "response_time": "840"
                },
                {
                    "country": "New Zealand",
                    "provider": "Kildy",
                    "bandwidth": "92",
                    "response_time": "1063"
                },
                {
                    "country": "Peru",
                    "provider": "Topolo",
                    "bandwidth": "2",
                    "response_time": "36"
                },
                {
                    "country": "Russian Federation",
                    "provider": "Topolo",
                    "bandwidth": "82",
                    "response_time": "830"
                },
                {
                    "country": "Saint Barthélemy",
                    "provider": "Kildy",
                    "bandwidth": "48",
                    "response_time": "667"
                },
                {
                    "country": "Spain",
                    "provider": "Topolo",
                    "bandwidth": "55",
                    "response_time": "790"
                },
                {
                    "country": "Switzerland",
                    "provider": "Topolo",
                    "bandwidth": "45",
                    "response_time": "1921"
                },
                {
                    "country": "Türkiye",
                    "provider": "Rond",
                    "bandwidth": "86",
                    "response_time": "1066"
                },
                {
                    "country": "United Kingdom of Great Britain and Northern Ireland",
                    "provider": "Topolo",
                    "bandwidth": "92",
                    "response_time": "1271"
                },
                {
                    "country": "United States of America",
                    "provider": "Rond",
                    "bandwidth": "50",
                    "response_time": "919"
                }
            ]
        ],
        "voice_call": [
            {
                "country": "RU",
                "bandwidth": "26",
                "response_time": "1371",
                "provider": "TransparentCalls",
                "connection_stability": 0.95,
                "ttfb": 1371,
                "voice_purity": 87,
                "median_of_call_time": 43
            },
            {
                "country": "US",
                "bandwidth": "78",
                "response_time": "1716",
                "provider": "E-Voice",
                "connection_stability": 0.89,
                "ttfb": 1716,
                "voice_purity": 27,
                "median_of_call_time": 7
            },
            {
                "country": "GB",
                "bandwidth": "21",
                "response_time": "1949",
                "provider": "TransparentCalls",
                "connection_stability": 0.95,
                "ttfb": 1949,
                "voice_purity": 43,
                "median_of_call_time": 56
            },
            {
                "country": "FR",
                "bandwidth": "12",
                "response_time": "1823",
                "provider": "TransparentCalls",
                "connection_stability": 0.93,
                "ttfb": 1823,
                "voice_purity": 44,
                "median_of_call_time": 18
            },
            {
                "country": "BL",
                "bandwidth": "91",
                "response_time": "1840",
                "provider": "E-Voice",
                "connection_stability": 0.61,
                "ttfb": 1840,
                "voice_purity": 78,
                "median_of_call_time": 14
            },
            {
                "country": "AT",
                "bandwidth": "44",
                "response_time": "171",
                "provider": "TransparentCalls",
                "connection_stability": 0.64,
                "ttfb": 171,
                "voice_purity": 43,
                "median_of_call_time": 30
            },
            {
                "country": "BG",
                "bandwidth": "39",
                "response_time": "164",
                "provider": "E-Voice",
                "connection_stability": 0.76,
                "ttfb": 164,
                "voice_purity": 34,
                "median_of_call_time": 4
            },
            {
                "country": "DK",
                "bandwidth": "11",
                "response_time": "1939",
                "provider": "JustPhone",
                "connection_stability": 0.73,
                "ttfb": 1939,
                "voice_purity": 33,
                "median_of_call_time": 48
            },
            {
                "country": "CA",
                "bandwidth": "57",
                "response_time": "797",
                "provider": "JustPhone",
                "connection_stability": 0.65,
                "ttfb": 797,
                "voice_purity": 18,
                "median_of_call_time": 3
            },
            {
                "country": "ES",
                "bandwidth": "83",
                "response_time": "668",
                "provider": "E-Voice",
                "connection_stability": 0.81,
                "ttfb": 668,
                "voice_purity": 43,
                "median_of_call_time": 42
            },
            {
                "country": "CH",
                "bandwidth": "50",
                "response_time": "1635",
                "provider": "JustPhone",
                "connection_stability": 0.84,
                "ttfb": 1635,
                "voice_purity": 76,
                "median_of_call_time": 25
            },
            {
                "country": "TR",
                "bandwidth": "40",
                "response_time": "157",
                "provider": "TransparentCalls",
                "connection_stability": 0.75,
                "ttfb": 157,
                "voice_purity": 35,
                "median_of_call_time": 56
            },
            {
                "country": "PE",
                "bandwidth": "47",
                "response_time": "145",
                "provider": "JustPhone",
                "connection_stability": 0.73,
                "ttfb": 145,
                "voice_purity": 52,
                "median_of_call_time": 46
            },
            {
                "country": "NZ",
                "bandwidth": "53",
                "response_time": "357",
                "provider": "JustPhone",
                "connection_stability": 0.74,
                "ttfb": 357,
                "voice_purity": 9,
                "median_of_call_time": 15
            },
            {
                "country": "MC",
                "bandwidth": "25",
                "response_time": "733",
                "provider": "E-Voice",
                "connection_stability": 0.9,
                "ttfb": 733,
                "voice_purity": 25,
                "median_of_call_time": 26
            }
        ],
        "email": {
            "AT": [
                [
                    {
                        "country": "AT",
                        "provider": "Gmail",
                        "delivery_time": 437
                    },
                    {
                        "country": "AT",
                        "provider": "Orange",
                        "delivery_time": 442
                    },
                    {
                        "country": "AT",
                        "provider": "RediffMail",
                        "delivery_time": 594
                    }
                ],
                [
                    {
                        "country": "AT",
                        "provider": "Yahoo",
                        "delivery_time": 29
                    },
                    {
                        "country": "AT",
                        "provider": "Live",
                        "delivery_time": 41
                    },
                    {
                        "country": "AT",
                        "provider": "Hotmail",
                        "delivery_time": 99
                    }
                ]
            ],
            "BG": [
                [
                    {
                        "country": "BG",
                        "provider": "MSN",
                        "delivery_time": 443
                    },
                    {
                        "country": "BG",
                        "provider": "RediffMail",
                        "delivery_time": 504
                    },
                    {
                        "country": "BG",
                        "provider": "Yahoo",
                        "delivery_time": 577
                    }
                ],
                [
                    {
                        "country": "BG",
                        "provider": "Hotmail",
                        "delivery_time": 129
                    },
                    {
                        "country": "BG",
                        "provider": "Yandex",
                        "delivery_time": 137
                    },
                    {
                        "country": "BG",
                        "provider": "Live",
                        "delivery_time": 166
                    }
                ]
            ],
            "BL": [
                [
                    {
                        "country": "BL",
                        "provider": "Hotmail",
                        "delivery_time": 472
                    },
                    {
                        "country": "BL",
                        "provider": "RediffMail",
                        "delivery_time": 517
                    },
                    {
                        "country": "BL",
                        "provider": "Live",
                        "delivery_time": 540
                    }
                ],
                [
                    {
                        "country": "BL",
                        "provider": "AOL",
                        "delivery_time": 28
                    },
                    {
                        "country": "BL",
                        "provider": "Mail.ru",
                        "delivery_time": 68
                    },
                    {
                        "country": "BL",
                        "provider": "MSN",
                        "delivery_time": 71
                    }
                ]
            ],
            "CA": [
                [
                    {
                        "country": "CA",
                        "provider": "AOL",
                        "delivery_time": 467
                    },
                    {
                        "country": "CA",
                        "provider": "MSN",
                        "delivery_time": 508
                    },
                    {
                        "country": "CA",
                        "provider": "Mail.ru",
                        "delivery_time": 554
                    }
                ],
                [
                    {
                        "country": "CA",
                        "provider": "Orange",
                        "delivery_time": 26
                    },
                    {
                        "country": "CA",
                        "provider": "GMX",
                        "delivery_time": 124
                    },
                    {
                        "country": "CA",
                        "provider": "Yahoo",
                        "delivery_time": 281
                    }
                ]
            ],
            "CH": [
                [
                    {
                        "country": "CH",
                        "provider": "Gmail",
                        "delivery_time": 443
                    },
                    {
                        "country": "CH",
                        "provider": "Orange",
                        "delivery_time": 519
                    },
                    {
                        "country": "CH",
                        "provider": "Yahoo",
                        "delivery_time": 596
                    }
                ],
                [
                    {
                        "country": "CH",
                        "provider": "Hotmail",
                        "delivery_time": 32
                    },
                    {
                        "country": "CH",
                        "provider": "Mail.ru",
                        "delivery_time": 42
                    },
                    {
                        "country": "CH",
                        "provider": "MSN",
                        "delivery_time": 49
                    }
                ]
            ],
            "DK": [
                [
                    {
                        "country": "DK",
                        "provider": "Comcast",
                        "delivery_time": 480
                    },
                    {
                        "country": "DK",
                        "provider": "Live",
                        "delivery_time": 533
                    },
                    {
                        "country": "DK",
                        "provider": "Yandex",
                        "delivery_time": 595
                    }
                ],
                [
                    {
                        "country": "DK",
                        "provider": "MSN",
                        "delivery_time": 93
                    },
                    {
                        "country": "DK",
                        "provider": "GMX",
                        "delivery_time": 158
                    },
                    {
                        "country": "DK",
                        "provider": "Yahoo",
                        "delivery_time": 160
                    }
                ]
            ],
            "ES": [
                [
                    {
                        "country": "ES",
                        "provider": "Comcast",
                        "delivery_time": 417
                    },
                    {
                        "country": "ES",
                        "provider": "Gmail",
                        "delivery_time": 479
                    },
                    {
                        "country": "ES",
                        "provider": "Live",
                        "delivery_time": 576
                    }
                ],
                [
                    {
                        "country": "ES",
                        "provider": "RediffMail",
                        "delivery_time": 14
                    },
                    {
                        "country": "ES",
                        "provider": "Hotmail",
                        "delivery_time": 45
                    },
                    {
                        "country": "ES",
                        "provider": "Yahoo",
                        "delivery_time": 185
                    }
                ]
            ],
            "FR": [
                [
                    {
                        "country": "FR",
                        "provider": "GMX",
                        "delivery_time": 412
                    },
                    {
                        "country": "FR",
                        "provider": "Gmail",
                        "delivery_time": 472
                    },
                    {
                        "country": "FR",
                        "provider": "Orange",
                        "delivery_time": 493
                    }
                ],
                [
                    {
                        "country": "FR",
                        "provider": "MSN",
                        "delivery_time": 51
                    },
                    {
                        "country": "FR",
                        "provider": "Yandex",
                        "delivery_time": 128
                    },
                    {
                        "country": "FR",
                        "provider": "Hotmail",
                        "delivery_time": 230
                    }
                ]
            ],
            "GB": [
                [
                    {
                        "country": "GB",
                        "provider": "RediffMail",
                        "delivery_time": 375
                    },
                    {
                        "country": "GB",
                        "provider": "MSN",
                        "delivery_time": 422
                    },
                    {
                        "country": "GB",
                        "provider": "Comcast",
                        "delivery_time": 514
                    }
                ],
                [
                    {
                        "country": "GB",
                        "provider": "Hotmail",
                        "delivery_time": 37
                    },
                    {
                        "country": "GB",
                        "provider": "Orange",
                        "delivery_time": 57
                    },
                    {
                        "country": "GB",
                        "provider": "GMX",
                        "delivery_time": 92
                    }
                ]
            ],
            "MC": [
                [
                    {
                        "country": "MC",
                        "provider": "Hotmail",
                        "delivery_time": 422
                    },
                    {
                        "country": "MC",
                        "provider": "Yahoo",
                        "delivery_time": 501
                    },
                    {
                        "country": "MC",
                        "provider": "MSN",
                        "delivery_time": 584
                    }
                ],
                [
                    {
                        "country": "MC",
                        "provider": "Mail.ru",
                        "delivery_time": 119
                    },
                    {
                        "country": "MC",
                        "provider": "Gmail",
                        "delivery_time": 163
                    },
                    {
                        "country": "MC",
                        "provider": "Yandex",
                        "delivery_time": 211
                    }
                ]
            ],
            "NZ": [
                [
                    {
                        "country": "NZ",
                        "provider": "Yahoo",
                        "delivery_time": 472
                    },
                    {
                        "country": "NZ",
                        "provider": "Comcast",
                        "delivery_time": 519
                    },
                    {
                        "country": "NZ",
                        "provider": "Gmail",
                        "delivery_time": 565
                    }
                ],
                [
                    {
                        "country": "NZ",
                        "provider": "RediffMail",
                        "delivery_time": 96
                    },
                    {
                        "country": "NZ",
                        "provider": "GMX",
                        "delivery_time": 110
                    },
                    {
                        "country": "NZ",
                        "provider": "Mail.ru",
                        "delivery_time": 139
                    }
                ]
            ],
            "PE": [
                [
                    {
                        "country": "PE",
                        "provider": "AOL",
                        "delivery_time": 422
                    },
                    {
                        "country": "PE",
                        "provider": "GMX",
                        "delivery_time": 502
                    },
                    {
                        "country": "PE",
                        "provider": "Gmail",
                        "delivery_time": 511
                    }
                ],
                [
                    {
                        "country": "PE",
                        "provider": "Hotmail",
                        "delivery_time": 58
                    },
                    {
                        "country": "PE",
                        "provider": "Live",
                        "delivery_time": 58
                    },
                    {
                        "country": "PE",
                        "provider": "Yahoo",
                        "delivery_time": 127
                    }
                ]
            ],
            "RU": [
                [
                    {
                        "country": "RU",
                        "provider": "Orange",
                        "delivery_time": 456
                    },
                    {
                        "country": "RU",
                        "provider": "Live",
                        "delivery_time": 515
                    },
                    {
                        "country": "RU",
                        "provider": "Mail.ru",
                        "delivery_time": 595
                    }
                ],
                [
                    {
                        "country": "RU",
                        "provider": "GMX",
                        "delivery_time": 130
                    },
                    {
                        "country": "RU",
                        "provider": "AOL",
                        "delivery_time": 163
                    },
                    {
                        "country": "RU",
                        "provider": "Hotmail",
                        "delivery_time": 199
                    }
                ]
            ],
            "TR": [
                [
                    {
                        "country": "TR",
                        "provider": "Mail.ru",
                        "delivery_time": 496
                    },
                    {
                        "country": "TR",
                        "provider": "Comcast",
                        "delivery_time": 581
                    },
                    {
                        "country": "TR",
                        "provider": "Yahoo",
                        "delivery_time": 593
                    }
                ],
                [
                    {
                        "country": "TR",
                        "provider": "Gmail",
                        "delivery_time": 2
                    },
                    {
                        "country": "TR",
                        "provider": "GMX",
                        "delivery_time": 22
                    },
                    {
                        "country": "TR",
                        "provider": "Hotmail",
                        "delivery_time": 54
                    }
                ]
            ],
            "US": [
                [
                    {
                        "country": "US",
                        "provider": "Live",
                        "delivery_time": 476
                    },
                    {
                        "country": "US",
                        "provider": "Mail.ru",
                        "delivery_time": 555
                    },
                    {
                        "country": "US",
                        "provider": "Yandex",
                        "delivery_time": 598
                    }
                ],
                [
                    {
                        "country": "US",
                        "provider": "GMX",
                        "delivery_time": 3
                    },
                    {
                        "country": "US",
                        "provider": "Hotmail",
                        "delivery_time": 13
                    },
                    {
                        "country": "US",
                        "provider": "Gmail",
                        "delivery_time": 54
                    }
                ]
            ]
        },
        "billing": {
            "create_customer": true,
            "purchase": true,
            "payout": false,
            "recurring": false,
            "fraud_control": true,
            "checkout_page": true
        },
        "support": [
            3,
            80
        ],
        "incident": [
            {
                "topic": "SMS delivery in EU",
                "status": "closed"
            },
            {
                "topic": "MMS connection stability",
                "status": "closed"
            },
            {
                "topic": "Voice call connection purity",
                "status": "closed"
            },
            {
                "topic": "Checkout page is down",
                "status": "closed"
            },
            {
                "topic": "Support overload",
                "status": "closed"
            },
            {
                "topic": "Buy phone number not working in US",
                "status": "closed"
            },
            {
                "topic": "API Slow latency",
                "status": "closed"
            }
        ]
    },
    "error": ""
}
```
