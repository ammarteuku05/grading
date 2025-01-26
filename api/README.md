# user

## Installation

Install the dependencies and devDependencies and start the server.

```sh
go mod tidy
```

## Migration 
```sh
make migrate-new // create new script for migrate 
make migration-up //up script 
```

## Run 
```sh
make run
```
## Build 
```sh
docker build -t interface-teacher-grading-api .
docker-compose up --build -d 
```

## List API
```sh
/v1/user/register [POST]
```
`Request :`
```sh
{
    "full_name":"ammar",
    "password":"12345678",
    "email":"student@mailnesia.com",
    "role" : "student"
}
```
`Response :`
```sh
{
    "code": 200,
    "message": "success",
    "data": null
}
```
```sh
/v1/user/login [POST]
```
`Request :`
```sh
{
    "password":"12345678",
    "email":"student@mailnesia.com",
}
```
`Response :`
```sh
{
    "code": 200,
    "message": "success",
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN0dWRlbnRAbWFpbG5lc2lhLmNvbSIsImZ1bGxfbmFtZSI6ImFtbWFyIiwiaWQiOiIzOTNkOWI5OS02YTlhLTQzZTMtOGFiMC0zZDkwMTRmZTRkMWEiLCJyb2xlIjoic3R1ZGVudCJ9.Arg1g6sKJrOduawiEZmG9o6TDz5e5Pz7WdNV74Hecdk"
}
```
```sh
/v1/assignment/create-assignment [POST]
```
`Header :`
```sh
{
  "Authorization": "Bearer <JWT-Token>"
}
```
`Request :`
```sh
{
    "subject":"Math",
    "tittle":"grammar",
    "content":"this content"
}
```
`Response :`
```sh
{
    "code": 200,
    "message": "success",
    "data": null
}
```
```sh
/v1/assignment/get-assignment [GET]
```
`Header :`
```sh
{
  "Authorization": "Bearer <JWT-Token>"
}
```
`Query Param :`
```sh
{
  "subject": "Math",
  "page":1,
  "per_page":1
}
```
`Response :`
```sh
{
    "code": 200,
    "message": "success",
    "data": {
        "page": 1,
        "per_page": 2,
        "total_page": 1,
        "total_data": 1,
        "items": [
            {
                "id": "6e3caf0c-7cb0-43c0-ab2e-79a63ebc1ebf",
                "subject": "English",
                "tittle": "grammar",
                "student_id": "45647a85-1db6-45c6-8b18-3b883e2d6b85",
                "content": "this content",
                "status": null,
                "CreatedAt": "2025-01-26T06:22:46.818029Z",
                "UpdatedAt": "2025-01-26T06:22:46.818029Z",
                "DeletedAt": null
            }
        ]
    }
}    
```
```sh
/v1/grade/get-grade [GET]
```
`Header :`
```sh
{
  "Authorization": "Bearer <JWT-Token>"
}
```
`Query Param :`
```sh
{
  "page":1,
  "per_page":1
}
```
`Response :`
```sh
{
    "code": 200,
    "message": "success",
    "data": {
        "page": 1,
        "per_page": 10,
        "total_page": 1,
        "total_data": 1,
        "items": [
            {
                "id": "68d96c1c-9e80-4be3-a42b-b2314b26dba3",
                "assignment_id": "636ff736-dbfc-4b5b-893c-fc75d1e69ec7",
                "teacher_id": "54fc1f51-d928-4cdf-a9a5-35adb3e3b520",
                "score": 90,
                "feedback": "test",
                "assignment": {
                    "id": "636ff736-dbfc-4b5b-893c-fc75d1e69ec7",
                    "subject": "Math",
                    "tittle": "grammar",
                    "student_id": "45647a85-1db6-45c6-8b18-3b883e2d6b85",
                    "content": "this content"
                }
            }
        ]
    }
}   
```
```sh
/v1/grade/create-grade [POST]
```
`Header :`
```sh
{
  "Authorization": "Bearer <JWT-Token>"
}
```
`Request :`
```sh
{
    "assignment_id":"636ff736-dbfc-4b5b-893c-fc75d1e69ec7",
	"score":90,
	"feedback":"test"
}
```
`Response :`
```sh
{
    "code": 200,
    "message": "success",
    "data": null
} 
```