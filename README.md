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