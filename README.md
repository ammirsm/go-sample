# Pilot project of Wealth Ethical

This project is now accessible from [heroku](https://wealth-ethical.herokuapp.com). 

This repository contains a golang backend for using in the pilot project of wealth ethical.
I've used [gorilla/mux](https://github.com/gorilla/mux) for router and [jinzhu/gorm](https://github.com/jinzhu/gorm) for my database models and ORM.

## Preparation

You can deploy this project with Godep :)

## Endpoints


| resource      | description                       |
|:--------------|:----------------------------------|
| `/cards`      | returns a list of cards (pass `from` and `limit` for pagination in query parameters)
| `/transactions/add_tag/`    | add tag to transaction
| `/transactions/delete_tag/{tag_id}` | delete a tag by its id |
| `/transactions/{account_id}/`      | returns a list of transactions that related to an account (pass `from` and `limit` for pagination of transactions in query parameters and you should pass `from_day` and `limit_day` for date filter.)|

