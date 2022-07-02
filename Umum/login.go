(POST) https://api-capstone-heroku.herokuapp.com/login

request 1:
{
    "email": "tia@gmail.com",
    "password": "tia12345"
}

respon 1:
{
    "code": 200,
    "id": 3,
    "level": "admin"
}

request 2:
{
    "email": "naufal@gmail.com",
    "password": "naufal123"
}

respon 2:
{
    "code": 200,
    "id": 5,
    "level": "dokter"
}

request 3:
{
    "email": "random",
    "password": "random"
}

respon 3:
{
    "message": "Email atau Password salah",
    "code": 400
}