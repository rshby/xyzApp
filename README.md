# xyz App
## how to run 
clone project ini kemudian buka menggunakan terminal/cmd. Arahkan ke direktori root foler. <br>
Selanjutnya ketikkan `docker compose up -d` <br>
Maka aplikasi akan secara otomatis menjalankan beberapa container yang diperlukan.

## Endpoint API
### List Of Endpoint
| Endpoint  | Method |
| ------------- |:------:|
| http://localhost:5005/v1/register|  POST  |
| http://localhost:5005/v1/tenor|  POST  |
| http://localhost:5005/v1/buy|  POST  |
| http://localhost:5005/v1/account|  POST  |
| http://localhost:5005/v1/login|  POST  |<br>

## Database
Database yang digunakan adalah MySQL yang otomatis jalan menggunakan container docker. <br>
Jika ingin membuka koneksi dengan database, dapat menggunakan configurasi berikut. <br>
```
user = root
password = root
host = localhost
port = 3306
```

## Tracing
Aplikasi ini juga mendukung tracing menggunakan Jaeger. <br>
Anda dapat membua dashboar jaeger dengan url `http://localhost:16686` <br>
Dengan menggunakan tracing ini kita akan lebih mudah mengetahui masing-masing proses yang jalan di apliasi ini. <br>
Harapannya penambahan tracing ini akan mempermudah kita dalam menemukan root cause ketika apliasi ini dirasa memiliki response time yang lama.<br>
[<img src="https://drive.google.com/uc?export=view&id=1O46mhOzEEtgjzJYrhU-396Gnb7KHcOlT" width="500"/>](https://drive.google.com/uc?export=view&id=1O46mhOzEEtgjzJYrhU-396Gnb7KHcOlT)

## Logging ELK
Aplikasi ini juga mendukung logging yang ditampilkan dari dahshboard Kibana.<br>
Anda dapat membuka dashboard Kibana dengan url berikut `http://localhost:5601` <br>
Dengan menggunakan filebeat-logstash-elasticsearch-kibana harapannya kita akan lebih mudah dalam melihat log untuk setiap request yang masuk ke aplikasi ini<br>
[<img src="https://drive.google.com/uc?export=view&id=1iwWL9SLZ7oKfUXZBT9MokHDxYaeRbUmK" width="500"/>](https://drive.google.com/uc?export=view&id=1O46mhOzEEtgjzJYrhU-396Gnb7KHcOlT)
