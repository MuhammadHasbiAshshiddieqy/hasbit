***TODO***
-	Add unit test (DONE)
-	Run `go test -v ./...` to test (DONE)
-	Lint using `golint {target}` (DONE)

***How to test***
-	Number_1 : Bisa dijalankan melalui `https://www.db-fiddle.com/`
-	Number_2 : Buat database (port 3306) `hasbit_external` dan `hasbit_user`, lalu jalankan semua service
-	Number_3 : Dapat dijalankan dengan `go test -v ./...` 
-	Number_4 : Dapat dijalankan dengan `go test -v ./...` 

***Flow Number_2***
Setelah menjalankan semua service :

-	Signup

URL (POST) : `http://localhost:8080/user/signup`

Body : request/signup.json

-	Login (Get token)

URL (POST) : `http://localhost:8080/user/login`

Body : request/login.json

-	Get Movies

URL (GET) : `http://localhost:8080/external/movies/{search}/{page}`

example : `http://localhost:8080/external/movies/Batman/2`

Header : `Authorization : {token dari endpoint /login}`

-	Get Movie by ID

URL (GET) : `http://localhost:8080/external/movie/{id}`

example : `http://localhost:8080/external/movie/tt0060153`

Header : `Authorization : {token dari endpoint /login}`


Contoh loadbalancer = `https://echo.labstack.com/middleware/proxy/` (hanya catatan pribadi)
