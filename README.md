hi, this is a test I did when applying for a job but because I have wifi problems I missed the deadline.
From here I learned that having more than 1 internet device is important.
if you want to try please move all folders to a new folder.

If you want try pleaae follow the step

## First Step

open your terminal

run this code :
```sh
go mod init golang-gorm
```

or run this code:

```sh
go mod init (name folder)
```

## Second Step

```sh
 go get -u github.com/gin-gonic/gin
```

## Third Step
```sh
 go get -u gorm.io/gorm
```

## Fourth Step
Run this code

```sh
go get gorm.io/driver/postgres
```

make new db and input your user, password, and dbname in file db.go in folder database
```sh
var (
	host     = "localhost"
	user     = "" //please input your user db
	password = "" //please input your pass db
	dBport   = "5432"
	dBname   = ""
	db       *gorm.DB
	err      error
)
```

Open the file db.go in folder database
search this code

```sh
//db.Debug().AutoMigrate()
```
uncomment this code because this code make table in your database

And fill your database user and password

## Testing

For the test in my rest-api, you can use this postman

```sh
https://api.postman.com/collections/13916221-79056d87-af9c-4b08-89f7-428ba37194d6?access_key=PMAT-01GN5F7A66PBCQGFXKY0XB1NP2
```

