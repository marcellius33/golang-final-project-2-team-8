# Hacktiv8 Golang Final Project 2 Team 8

## Link Deploy API

https://golang-final-project-2-team-8-production.up.railway.app

## Link Postman Documentation

https://documenter.getpostman.com/view/21171989/2s8YmGW75y

## Anggota
1. Mochamad Suhri Ainur Rifky (GLNG-KS04-001)
2. Raden Muhammad Yudie Sanjaya (GLNG-KS04-016) :x:
3. Varrel Marcellius (GLNG-KS04-021)

## Pembagian Tugas
1. Mochamad Suhri Ainur Rifky (GLNG-KS04-001)
- Setup project
- API User
- Postman (Collection, Environment, Documentation)
- Unit Test

2. Raden Muhammad Yudie Sanjaya (GLNG-KS04-016) :x:


3. Varrel Marcellius (GLNG-KS04-021)
- API Photo
- API Comment
- API Social Media
- Readme
- Deploy

# Cara Install
1. run `docker compose up` untuk menjalankan database
2. run `go run main.go` untuk menjalankan applikasi

Nama File Postman: -

# List Route
## Users
- `POST` - `/users/register`. Digunakan untuk registrasi pengguna
- `POST` - `/users/login`. Digunakan untuk login pengguna
- `PUT` - `/users/:userId`. Digunakan untuk memperbaharui detail pengguna
- `DELETE` - `/users/:userId`. Digunakan untuk menghapus pengguna

## Photos
- `GET` - `/photos`. Digunakan untuk mengambil list photo dari pengguna 
- `POST` - `/photos`. Digunakan untuk membuat photo baru
- `PUT` - `/photos/:photoId`. Digunakan untuk memperbaharui photo
- `DELETE` - `/photos/:photoId`. Digunakan untuk menghapus photo

## Comment
- `GET` - `/comments`. Digunakan untuk mengambil list comment dari pengguna
- `POST` - `/comments`. Digunakan untuk membuat comment baru terhadap post
- `PUT` - `/comments/:commentId`. Digunakan untuk memperbaharui comment
- `DELETE` - `/comments/:commentId`. Digunakan untuk menghapus comment

## Social Media
- `GET` - `/socialmedias`. Digunakan untuk mengambil list social media dari pengguna
- `POST` - `/socialmedias`. Digunakan untuk membuat social media baru terhadap post
- `PUT` - `/socialmedias/:socialMediaId`. Digunakan untuk memperbaharui social media
- `DELETE` - `/socialmedias/:socialMediaId`. Digunakan untuk menghapus social media