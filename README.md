# Introduction

Welcome to the whistestop tour of Go! This session is an interactive workshop giving practical experience building a web based service in go. It'll be divided into 4 parts, each building on the last to deliver a weather application (with a twist), and touch on some of the aspects of daily programming in Go.

Estimated Duration: 1 hours

### Session Overview
The session is divided into 4 parts, each building upon the last, and covers the following:

**Part 1:** Setting up, running and building a Go project and modules.

**Part 2:** Creating and packaging an integration with 3rd party weather API to fetch data.

**Part 3:** Implementing our business rules, and testing the implementation to reduce bugs.

**Part 4:** Defining a http server to expose our functionality as an API.

---
### Getting Setup
To take part in this session, you'll need atleast the following setup:
1) An installation of Go - [Installation instructions](https://golang.org/doc/install)
2) Your favourite editor to write some Go - [VSCode](https://code.visualstudio.com/) is a great tool for the job if you don't already have a favourite.
3) A tool like curl or [Postman](https://www.postman.com/) to test the web service.

---
### The Specification
We will be building an API to inform users what type of attire they should wear for the day, based on the weather in their area.
For example, if it's raining in London, we should advise them to wear waterproofs and take an umbrella ☔.
The API must be made available over http, with the endpoint defined below. 

**Rules**:
1) If it's raining, we should always suggest an umbrella.
2) If it's raining and the temp is below **20 degrees** we should suggest a waterproof and windproof jacket, and trousers
3) If it's raining and the temp is above **20 degrees**, we should suggest a waterproof jacket, and shorts.

**API Requirements**:
1) `GET /v1/attire?location=XYZ` - This endpoint should return a decision to the user informing them on what attire they should wear for the day ahead. 
   ```json
   200 OKAY
   {
       "attire": {
           "jacket": {
               "waterproof": true,
               "windproof": true
           },
           "pants": "trousers",
           "umbrella": true
       }
   }
   400 BAD REQUEST{}
   ```
---

### Weather API
We'll be using an external API to fetch weather information that we will use for our service, below is the specification:

#### Fetch current weather information for a location
`GET v1/weather?location=XYZ`

Response:
```json
200 OKAY
{
    "main": "Clear",
    "temp_min": 12.1,
    "temp_max": 14.1,
    "wind_speed": 4,
}
400 BAD REQUEST
{}
401 UNAUTHORIZED
{}
500 INTERNAL SERVER ERROR
{}
```

The possible `main` values are:
```
Rain
Heavy Rain
Sun
Cold
Overcast
```

A URL will be shared during the workshop, enabling you to call the API over the internet

**DISCLAIMER**
The external API you will be using was built for the purpose of this workshop, and is not accurate data :)

---
### Futher Reading
Some useful sites that offer great overview and detail of the Go language.
1) [Tour of Go](https://tour.golang.org/welcome/1)
2) [Effective Go](https://golang.org/doc/effective_go.html)
3) [Go By Example](https://gobyexample.com/)

---
### License
Copyright © Matthew Nolf 2020 