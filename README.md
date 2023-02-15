<pre style="font-size: 1.4vw;">
<p align="center">
             _                                     _ _                                    
            | |                                   | (_)               _              _    
  ____  ___ | | ____ ____   ____     ____ ___   _ | |_ ____   ____   | |_  ____  ___| |_  
 / _  |/ _ \| |/ _  |  _ \ / _  |   / ___) _ \ / || | |  _ \ / _  |  |  _)/ _  )/___)  _) 
( ( | | |_| | ( ( | | | | ( ( | |  ( (__| |_| ( (_| | | | | ( ( | |  | |_( (/ /|___ | |__ 
 \_|| |\___/|_|\_||_|_| |_|\_|| |   \____)___/ \____|_|_| |_|\_|| |   \___)____|___/ \___)
(_____|                   (_____|                           (_____|                       
</p>
</pre>
<p align="center">
<a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg">
</a>
<a href="/LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-green.svg">
</a>
</p>
<p align="center">
RESTful API of <b>Golang Coding Test</b>
</p>

# Golang Conding Test of API Guide

## 🔀 Compatible Route Endpoint
| NO | Use                           | Endpoint               | Example                     | Action
|----|-------------------------------|------------------------|-----------------------------|------------
| 1  | Hello                         | /              | http://localhost:3003/              | GET
| 2  | Palindrom                     | /palindrom     | http://localhost:3003/palindrom     | GET
| 3  | Get language point 2          | /language      | http://localhost:3003/language      | GET
| 4  | Get All language              | /languages     | http://localhost:3003/languages     | POST
| 5  | Get language by id            | /language/:id  | http://localhost:3003/language/:id  | GET
| 6  | Create language               | /language      | http://localhost:3003/language      | POST
| 7  | Get language by id            | /language/:id  | http://localhost:3003/language/:id  | GET
| 8  | Update language by id         | /language/:id  | http://localhost:3003/language/:id  | UPDATE
| 9  | Delete language by id         | /language/:id  | http://localhost:3003/language/:id  | DELETE

---

## 📖 Compatible JSON Payload Golang Conding Test of API
This is the JSON payload that's sended to Golang Conding Test of API

### 💲 Create Language JSON Payload (Raw JSON)
```js
[
    {
        "appeared": 2009,
        "language": "Go",
        "created": [
            "Robert Griesemer",
            "Rob Pike",
            "Ken Thompson"
        ],
        "functional": true,
        "object-oriented": false,
        "relation": {
            "influenced-by": [
                "C",
                "Python",
                "Pascal",
                "Smalltalk",
                "Modula"
            ],
            "influences": [
                "Odin",
                "Crystal",
                "Zig"
            ]
        }
    },
    {
        "appeared": 1991,
        "language": "Python",
        "created": [
            "Guido van Rossum"
        ],
        "functional": true,
        "object-oriented": true,
        "relation": {
            "influenced-by": [
                "C",
                "C++",
                "ALGOL 68",
                "Ada",
                "Haskell",
                "Java",
                "Lisp"
            ],
            "influences": [
                "Go",
                "CoffeeScript",
                "JavaScript",
                "Julia",
                "Ruby",
                "Swift"
            ]
        }
    }
]
```

### 💸 Get Language by ID
```js
http://localhost:3003/language/1 => GET
```
### 💸 Update Language JSON Payload (Raw JSON)

```js
http://localhost:3003/language/1 => UPDATE

{
    "language": "Laravel 10",
    "appeared": 2009,
    "created": [
        "Robert Griesemer",
        "Rob Pike",
        "Ken Thompson"
    ],
    "functional": true,
    "object-oriented": false,
    "relation": {
        "influenced-by": [
            "C",
            "Python",
            "Pascal",
            "Smalltalk",
            "Modula"
        ],
        "influences": [
            "Odin",
            "Crystal",
            "Zig"
        ]
    }
}
```

### 💸 Delete Language by ID
```js
http://localhost:3003/language/1  => DELETE 
```