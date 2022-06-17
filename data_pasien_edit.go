(PUT) https://api-capstone-heroku.herokuapp.com/auth/edit_data_pasien/id

request: semua Rest API yang mengandung "/auth" harus menggunakan token JWT untuk mengaksesnya
{
    "nik":"330707880",
    "nama":"Iqbal Ramadhan",
    "alamat":"Sulawesi Tenggara",
    "jenis_kelamin":"L",
    "nomer_telfon":"083716281911",
    "tempat_lahir":"London",
    "tanggal_lahir":"2002-06-18"
}

respon:
{
    "data": {
        "id": 6,
        "nik": "330707880",
        "nama": "Iqbal Ramadhan",
        "alamat": "Sulawesi Tenggara",
        "jenis_kelamin": "L",
        "no_hp": "083716281911",
        "tempat_lahir": "London",
        "tanggal_lahir": "2002-06-18"
    },
    "status": "berhasil mengedit data pasien",
    "userID": "candra@gmail.com"
}