# go-restful-api

## Overview
A restful API built in Golang with Clean Architecture. Study case for a phone service center. Here are the features that we want to achieve:
- Customer can input their order based on available service
- Customer can see their service cost, assigned technician, and when to pick their device
- System saved the order and assign the order to the correct technician
- Technician can see their customer line

API contract can be found [HERE](https://docs.google.com/spreadsheets/d/1UeOR79LOUlD5d7G6OtC5wtZGF0z_rmlUb75qXIZkWeE/edit?usp=sharing)

## Project Status
Not Finished Yet
To Do List:
- [x] Create boilerplate for the repo
- [x] Write sql migrations
- [x] Create models
- [] Create Domain and Usecases
-- [In Progress] Customer
-- [] Order
-- [] Services
-- [] Technician

## Technology
Here we will used
- Golang
- Postgres on heroku
- Redis
