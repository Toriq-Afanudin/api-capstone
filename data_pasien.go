(GET) https://api-capstone-heroku.herokuapp.com/auth/data_pasien

request: membutuhkan token jwt

respon:
{
    "data": [
        {
            "Id": 1,
            "Nik": "330707876",
            "Nama": "Thoriq Afanudin",
            "Alamat": "Demak",
            "Jenis_kelamin": "L",
            "Jenis_penyakit": "Magh",
            "No_hp": "082122214133",
            "Tempat_lahir": "Wonosobo",
            "Tanggal_lahir": "1999-12-18"
        },
        {
            "Id": 2,
            "Nik": "330707877",
            "Nama": "Ayu Wahyuni",
            "Alamat": "Wonosobo",
            "Jenis_kelamin": "P",
            "Jenis_penyakit": "Vertigo",
            "No_hp": "087658291008",
            "Tempat_lahir": "Demak",
            "Tanggal_lahir": "2000-06-11"
        },
        {
            "Id": 3,
            "Nik": "330707878",
            "Nama": "Trisna Candra",
            "Alamat": "Malang",
            "Jenis_kelamin": "L",
            "Jenis_penyakit": "Sakit gigi",
            "No_hp": "087658291555",
            "Tempat_lahir": "Malang",
            "Tanggal_lahir": "1998-11-21"
        },
        {
            "Id": 4,
            "Nik": "330707879",
            "Nama": "Dhiya Aisy",
            "Alamat": "Wonogiri",
            "Jenis_kelamin": "P",
            "Jenis_penyakit": "Patah hati",
            "No_hp": "087658291222",
            "Tempat_lahir": "Indramayu",
            "Tanggal_lahir": "1999-06-16"
        }
    ],
    "status": "berhasil menampilkan data pasien",
    "userID": "candra@gmail.com"
}