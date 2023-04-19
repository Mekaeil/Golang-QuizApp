# Golang QuizApp

Creating Quiz App application with Golang based on best practices and DDD structure.

## DB Diagram
You can see and find the [diagram in draw.io link](https://drive.google.com/file/d/1MqIh5MLh9cfK_rP3R-UUkcCde3Njt3ZK/view?usp=sharing), and also you can see the image of the DB diagram.

<p style="width: 30em;text-align:left;">

![image](https://github.com/Mekaeil/QuizApp/blob/master/public/images/QuizApp_diagram.png)

</p>

## QuizApp Logic & Features

In this small application we want to create the below features:

- question CRUD
    - question types [select, multi-select, text]
- question replies CRUD and assign to questions
- quiz-box to make a group of questions
- scheduled quiz-box to make it easy schedule a group of questions to send to users
    - we can define reply_type to have several type of replies
        - realtime
        - email
        - custom_correction_required
        - end_of_exam
- tag to create a label for different type of entities
- quiz-replier-activity: we will log user's answer 

# Installing the project
If you are not using Makefile just install it, and then follow below steps to install and using other features as well

## Prerequisites

#### Docker setup
- [macOS](https://docs.docker.com/desktop/mac/install/) 
- [Linux](https://runnable.com/docker/install-docker-on-linux) 
- [Windows](https://docs.docker.com/desktop/windows/install/)

#### Brew 
- [Install Homebrew - Mac](https://brew.sh/)
- [Brew Make](https://formulae.brew.sh/formula/make)
- [Learn Makefiles](https://makefiletutorial.com/)

## Installation Steps
``` 
// Install The Project
make install    

// Keep Watching the changes
make watch
```


# Packages
- [GORM](https://gorm.io/docs/) - DataBase ORM
  - [MySql Driver](https://gorm.io/docs/connecting_to_the_database.html#mysql) - MySql DataBase ORM
- [Echo](https://echo.labstack.com/guide) - Framework
- [Validator](https://github.com/go-playground/validator) - Data Input Validator
- [Reflex](https://github.com/cespare/reflex) - Reflex Live Reloading
- [GoDotEnv](https://github.com/joho/godotenv) - Loading .env file
