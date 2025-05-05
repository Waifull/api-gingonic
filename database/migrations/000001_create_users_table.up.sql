CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    address VARCHAR(255),
    email VARCHAR(255),
    born_date TIMESTAMP
);





--command untuk membuat file migrasi baru
--migrate create -ext sql -dir db/migrations -seq create_users_table
--seq untuk penomoran migrasi

--command untuk menjalankan semua migrasi up
--go-gin-gonic merupakan nama database
--migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/go-gin-gonic" -path database/migrations up