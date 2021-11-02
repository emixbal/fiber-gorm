# Go Fiber Boilerplate with Gorm ORM
This boilerplate app is using Go version 1.17 because I think for now this is the most updated release. 

## Installation

pastikan pc anda sudah terinstall go, pda envoirement kami menggunakan go versi 1.17:

```
$ git clone https://github.com/emixbal/fiber-gorm.git

$ cd fiber-gorm

$ go get
```

## Run Development Mode

- buat file denga nama .env pada root directory
- anda dapat menyesuaikan isi file .env dengan mencontoh .env.example
- Pada root directory, jalankan perintah ini
    ```
    $ go run main.go
    ```
- pada client requester (mis:postman) buka alamat localhost:3000/books