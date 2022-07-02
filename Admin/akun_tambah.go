(POST) https://api-capstone-heroku.herokuapp.com/admin/akun_tambah

request 1:
{
    "nickname":"bahtiar",
    "email":"bahtiar@gmail.com",
    "password":"bahtiar123",
    "level":"perawat"
}

respon 1:
{
    "code": 200,
    "data": {
        "nickname": "bahtiar",
        "email": "bahtiar@gmail.com",
        "password": "bahtiar123",
        "level": "perawat"
    },
    "message": "Perawat bahtiar diharuskan melengkapi data sendiri"
}

request 2:
{
    "nickname":"mansur",
    "email":"mansur@gmail.com",
    "password":"mansur123",
    "level":"dokter"
}

respon 2:
{
    "code": 200,
    "data": {
        "nickname": "mansur",
        "email": "mansur@gmail.com",
        "password": "mansur123",
        "level": "dokter"
    },
    "message": "Dokter mansur diharuskan melengkapi data sendiri"
}

request 3:
{
    "nickname":"dias",
    "email":"dias@gmail.com",
    "password":"dias1234",
    "level":"admin"
}

respon 3:
{
    "code": 200,
    "data": {
        "nickname": "dias",
        "email": "dias@gmail.com",
        "password": "dias1234",
        "level": "admin"
    },
    "message": "dias resmi menjadi admin baru"
}