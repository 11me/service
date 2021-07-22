package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	var nc *nats.Conn
	var err error
	message := []byte(`{
	  "order_uid": "afc4e94948fe433f8b55b68a0f169e56",
	  "entry": "WBIL2",
	  "internal_signature": "8062ffbe87d2446dbc98b05a0e4e07bd53565c0b2fd90c6c6e6a96594de61433",
	  "payment": {
	    "transaction": "afc4e94948fe433f8b55b68a0f169e56",
	    "currency": "USD",
	    "provider": "WbPay",
	    "amount": 3107,
	    "payment_dt": 1614540555,
	    "bank": "sber",
	    "delivery_cost": 0,
	    "goods_totalNOT": 3107
	  },
	  "items": [
	    {
	      "chrt_id": 11241243,
	      "price": 300,
	      "rid": "7ed1002bcbc14fa594d48b422ef6980d",
	      "name": "Сказки малышам. (библиотека детского сада).",
	      "sale": 45,
	      "size": "",
	      "total_price": 164,
	      "nm_id": 2786861,
	      "brand": "Умка"
	    },
	    {
	      "chrt_id": 29535331,
	      "price": 236,
	      "rid": "d7a9cd9d6cc04b48b50e00b9f4cd4c9a",
	      "name": "Лопушок. Сказки для малышей",
	      "sale": 28,
	      "size": "",
	      "total_price": 169,
	      "nm_id": 8771153,
	      "brand": "Амрита"
	    },
	    {
	      "chrt_id": 34856620,
	      "price": 259,
	      "rid": "758de71d0f764554839c53b4bb062f67",
	      "name": "Иван-крестьянский сын и Чудо-юдо",
	      "sale": 45,
	      "size": "",
	      "total_price": 142,
	      "nm_id": 10881092,
	      "brand": "Издательство Речь"
	    },
	    {
	      "chrt_id": 39310989,
	      "price": 316,
	      "rid": "69657189a2864ffbb0acb7adeaaf4705",
	      "name": "ЛЕВ ТОЛСТОЙ. РАССКАЗЫ, СКАЗКИ И БЫЛИ  \"ДЮЖИНА\"",
	      "sale": 41,
	      "size": "",
	      "total_price": 184,
	      "nm_id": 12759167,
	      "brand": "Проф-Пресс"
	    },
	    {
	      "chrt_id": 39310990,
	      "price": 303,
	      "rid": "4cdc75623c4c4e65969d2635f6f55598",
	      "name": "МОИ ПЕРВЫЕ СКАЗКИ",
	      "sale": 48,
	      "size": "",
	      "total_price": 157,
	      "nm_id": 12759168,
	      "brand": "Проф-Пресс"
	    },
	    {
	      "chrt_id": 39310991,
	      "price": 316,
	      "rid": "4e09c06d863844bab9f0eec9844d5de8",
	      "name": "РУССКИЕ СКАЗКИ МАЛЫШАМ  \"ДЮЖИНА\"",
	      "sale": 55,
	      "size": "",
	      "total_price": 142,
	      "nm_id": 12759169,
	      "brand": "Проф-Пресс"
	    },
	    {
	      "chrt_id": 39819896,
	      "price": 314,
	      "rid": "5787871dc18746ddbc97cbcb974fc470",
	      "name": "Записная книжка",
	      "sale": 51,
	      "size": "",
	      "total_price": 152,
	      "nm_id": 12993881,
	      "brand": "СИМА-ЛЕНД"
	    },
	    {
	      "chrt_id": 41096819,
	      "price": 886,
	      "rid": "a71041ee3c98464293149b1b3882858e",
	      "name": "Ночная сорочка",
	      "sale": 60,
	      "size": "60",
	      "total_price": 350,
	      "nm_id": 13587977,
	      "brand": "Alena Alenkina"
	    },
	    {
	      "chrt_id": 42079863,
	      "price": 179,
	      "rid": "521f177ea92e4df0a3b24a0d0a0990c3",
	      "name": "Федорино горе",
	      "sale": 20,
	      "size": "",
	      "total_price": 142,
	      "nm_id": 14065528,
	      "brand": "Школьная Книга"
	    },
	    {
	      "chrt_id": 42079877,
	      "price": 179,
	      "rid": "e7fb76e9e13641dc96181b4421c0516f",
	      "name": "Бармалей",
	      "sale": 20,
	      "size": "",
	      "total_price": 142,
	      "nm_id": 14065542,
	      "brand": "Школьная Книга"
	    },
	    {
	      "chrt_id": 42079888,
	      "price": 179,
	      "rid": "0f24b387352d441fa7bc3e3b028785d9",
	      "name": "Снегурочка",
	      "sale": 20,
	      "size": "",
	      "total_price": 142,
	      "nm_id": 14065553,
	      "brand": "Школьная Книга"
	    },
	    {
	      "chrt_id": 44683515,
	      "price": 179,
	      "rid": "d12cd58fca9a46f593cebda3c2ed087b",
	      "name": "Лиса и журавль. Журавль и цапля. ИГРАЕМ В СКАЗКУ. ТЕАТРАЛИЗАЦИЯ СКАЗОК с игровыми полями",
	      "sale": 20,
	      "size": "",
	      "total_price": 142,
	      "nm_id": 15346238,
	      "brand": "Школьная Книга"
	    },
	    {
	      "chrt_id": 44996935,
	      "price": 236,
	      "rid": "62abdf5eb55f4eafb0eeb02b813d8c7f",
	      "name": "Наушники",
	      "sale": 30,
	      "size": "",
	      "total_price": 164,
	      "nm_id": 15503637,
	      "brand": "Dream Tech"
	    },
	    {
	      "chrt_id": 45726498,
	      "price": 159,
	      "rid": "dd7a0f9a64d547978829cbd471d0f267",
	      "name": "\"КОЛОБОК\" КНИЖКА-ГАРМОШКА",
	      "sale": 36,
	      "size": "",
	      "total_price": 101,
	      "nm_id": 15889643,
	      "brand": "Проф-Пресс"
	    },
	    {
	      "chrt_id": 47551424,
	      "price": 1886,
	      "rid": "555c32eae827422bbb9bb377e1b2854b",
	      "name": "Тапочки",
	      "sale": 56,
	      "size": "43-44",
	      "total_price": 814,
	      "nm_id": 16962802,
	      "brand": "Naturella&Home"
	    }
	  ],
	  "locale": "ru",
	  "customer_id": "5ea488619943420eaefcbcc402eb8ddc",
	  "track_number": "WBIL2817015795SL",
	  "delivery_service": "meest",
	  "shardkey": "2",
	  "sm_id": 33
	}`)

	nc, err = nats.Connect("localhost")
	if err != nil {
		fmt.Println("Error occured: ", err)
	}

	nc.Publish("foo", []byte(message))

	time.Sleep(2 * time.Second)
}
