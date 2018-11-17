# RecipeWalkthrough
[![Build Status](https://travis-ci.org/Smart-Chef/trigger-queue.svg?branch=master)](https://travis-ci.org/Smart-Chef/recipe-walkthrough)
[![codecov](https://codecov.io/gh/Smart-Chef/trigger-queue/branch/master/graph/badge.svg)](https://codecov.io/gh/Smart-Chef/recipe-walkthrough)

The recipe walkthrough has two large functionalities. The first one is the initialization of a recipe. When a recipe along with all of it’s steps are triggers are sent to this service, it will first store this data in memory and also initialize the current state. The state has data about what step we are currently on. After storing this data, the triggers and the actions for the first step will then be sent over and registered in the trigger queue. 

The other functionality of this service is to either change the current step or the current recipe. If the action is to change the step (i.e. an action from a trigger is trying to increment to the next step) or recipe, both will first will ensure that all triggers for the currently active step have been removed from the trigger queue. If it a new step, the new step will be initialized. If the step is to change recipes, it will now enter the initialize recipe workflow.

## CLI
Use `-h` to see help

```shell
Usage of ./recipe-walkthrough:
  -setup
    	Run setup (includes creating db schema)
  -verbose
    	Display verbose logs
  -walk
    	Walk the routes
```

## Setup

Ensure you have a go version compatible with modules

```bash
git clone https://github.com/Smart-Chef/recipe-walkthrough # could also probably go get
cd recipe-walkthrough
go build
./recipe-walkthrough
```

## Loading data

You can load the DB schema with either of the two commands

```shell
./recipe-walkthrough -setup
# or
sqlite3 smart-chef.db < setup.sql
```

SQL scripts for mock data is located in `data`.
You can load them into the DB with `cat data/* | sqlite3 smart-chef.db`

## Dev Usage

The API docs can be found at [PostmanDocs](https://documenter.getpostman.com/view/1907478/RzZ6GzdW)

This service interacts with the [trigger-queue](https://github.com/Smart-Chef/trigger-queue)
and has the actual Recipe Storage with the DB as well as the main utility for driving the state
machine of the recipe from start to end.

Because this is a CGO enabled package used in this project you are required to set the environment variable
`CGO_ENABLED=1` and have a gcc compile present within your path