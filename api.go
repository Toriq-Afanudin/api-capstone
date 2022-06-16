API yang sudah bisa digunakan:
(GET) https://api-capstone-heroku.herokuapp.com/login
(GET) https://api-capstone-heroku.herokuapp.com/data_pasien
(POST) https://api-capstone-heroku.herokuapp.com/tambah_akun



PETUNJUK DEMO:

(GET) https://api-capstone-heroku.herokuapp.com/login

request 1:
{
    "Email":"thoriqafan03@gmail.com",
    "Password":"thoriq123"
}

respon 1:
{
    "email": "thoriqafan03@gmail.com",
    "level": "admin",
    "status": "login berhasil"
}

request 2:
{
    "Email":"candra@gmail.com",
    "Password":"candra123"
}

respon 2:
{
    "email": "candra@gmail.com",
    "level": "dokter",
    "status": "login berhasil"
}

request 3:
{
    "Email":"random",
    "Password":"random"
}

respon 3:
{
    "status": "email atau password salah"
}



(GET) https://api-capstone-heroku.herokuapp.com/data_pasien

respon:
{
    "akun": [
        {
            "Id": 1,
            "Nik": "330707876",
            "Nama": "Thoriq Afanudin",
            "Alamat": "Demak",
            "Jenis_kelamin": "L",
            "Jenis_penyakit": "Magh"
        },
        {
            "Id": 2,
            "Nik": "330707877",
            "Nama": "Ayu Wahyuni",
            "Alamat": "Wonosobo",
            "Jenis_kelamin": "P",
            "Jenis_penyakit": "Vertigo"
        },
        {
            "Id": 3,
            "Nik": "330707878",
            "Nama": "Trisna Candra",
            "Alamat": "Malang",
            "Jenis_kelamin": "L",
            "Jenis_penyakit": "Sakit gigi"
        },
        {
            "Id": 4,
            "Nik": "330707879",
            "Nama": "Dhiya Aisy",
            "Alamat": "Wonogiri",
            "Jenis_kelamin": "P",
            "Jenis_penyakit": "Patah hati"
        }
    ],
    "status": "berhasil menampilkan data pasien"
}



(POST) https://api-capstone-heroku.herokuapp.com/tambah_akun
request:
{
    "nickname":"fatoni",
    "email":"fathan@gmail.com",
    "password":"fathan123",
    "level":"perawat"
}

respon:
{
    "akun": "fatoni",
    "status": "akun berhasil di tambahkan"
}