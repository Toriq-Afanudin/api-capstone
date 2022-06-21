(POST) https://api-capstone-heroku.herokuapp.com/akun_tambah

request 1: menambahkan akun dokter
{
    "nickname":"syarif",
    "email":"syarif@gmail.com",
    "password":"syarif123",
    "level":"dokter"
}

respon 1:
{
    "data": "syarif",
    "message": "Lengkapi data dokter syarif",
    "status": "Berhasil",
    "user": null
}

request 2: menambahkan akun perawat
{
    "nickname":"ayu",
    "email":"ayu@gmail.com",
    "password":"ayu12345",
    "level":"perawat"
}

respon 2:
{
    "data": "ayu",
    "message": "Lengkapi data perawat ayu",
    "status": "Berhasil",
    "user": null
}

request 3: menambahkan akun admin
{
    "nickname":"yusuf",
    "email":"yusuf@gmail.com",
    "password":"yusuf123",
    "level":"admin"
}

respon 3:
{
    "data": "yusuf",
    "message": "",
    "status": "Berhasil",
    "user": null
}

request 4: data tidak boleh ada yang kosong
{
    "nickname":"",
    "email":"yusuf@gmail.com",
    "password":"yusuf123",
    "level":"admin"
}

respon 4:
{
    "message": "Data tidak boleh ada yang kosong.",
    "status": "Error"
}

request 5: pilih pekerjaan sebagai dokter, perawat, atau admin
{
    "nickname":"bahtiar",
    "email":"bahtiar@gmail.com",
    "password":"bahtiar123",
    "level":"office boy"
}

respon 5:
{
    "message": "Pilih jenis pekerjaan sebagai admin, dokter, atau perawat.",
    "status": "Error"
}