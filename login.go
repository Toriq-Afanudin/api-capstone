(POST) https://api-capstone-heroku.herokuapp.com/login

request 1:
{
    "email":"candra@gmail.com",
    "password":"candra123"
}

respon 1:
{
    "code": 200,
    "expire": "2022-06-17T18:09:43Z",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTU0ODkzODMsImlkIjoiY2FuZHJhQGdtYWlsLmNvbSIsIm9yaWdfaWF0IjoxNjU1NDg1NzgzfQ.YWNL2SxBgMq5zAAvj3GKOjaWnpfXX7dd64dTKGW1Dd0"
}

request 2:
{
    "email":"random",
    "password":"random"
}

respon 2:
{
    "code": 401,
    "message": "incorrect Username or Password"
}