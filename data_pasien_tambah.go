(POST) https://api-capstone-heroku.herokuapp.com/auth/tambah_data_pasien

request 1: membutuhkan token jwt
{
    "nik":"330707880",
    "nama":"Iqbal Ramadhan",
    "alamat":"Nusa Tenggara",
    "jenis_kelamin":"L",
    "nomer_telfon":"083716281911",
    "tempat_lahir":"Nusa Tenggara",
    "tanggal_lahir":"2002-06-18"
}

respon 1:
{
    "data": {
        "id": 6,
        "nik": "330707880",
        "nama": "Iqbal Ramadhan",
        "alamat": "Nusa Tenggara",
        "jenis_kelamin": "L",
        "no_hp": "083716281911",
        "tempat_lahir": "Nusa Tenggara",
        "tanggal_lahir": "2002-06-18"
    },
    "status": "berhasil menambahkan data pasien",
    "userID": "candra@gmail.com"
}

request 2:
{
    "nik":"330707880",
    "nama":"",
    "alamat":"Nusa Tenggara",
    "jenis_kelamin":"L",
    "nomer_telfon":"083716281911",
    "tempat_lahir":"",
    "tanggal_lahir":"2002-06-18"
}

respon 2:
{
    "status": "gagal menambahkan, tidak boleh ada data yang kosong"
}

request 3:
{
    "nik":"330707880",
    "nama":"Andi",
    "alamat":"Nusa Tenggara",
    "jenis_kelamin":"X",
    "nomer_telfon":"083716281911",
    "tempat_lahir":"Germany",
    "tanggal_lahir":"2002-06-18"
}

respon 3:
{
    "status": "gagal menambahkan, jenis kelamin harus di isi dengan P atau L"
}