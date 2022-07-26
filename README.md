# quiz_master
Quiz Master
================
It is a CRUD CLI implementation that simulates an quiz masterdata. It uses Mysql to store the data, and all CRUD commands and others can be implemented dinamically without the need to restart the main program to load it.

Required Install
================
```install docker```

How to build?
================
You can simple:
```bin/setup```

How to run?
================
Proceed using ```bin/quizmaster```

type `help`  to list out the command list with descriptions

List of commands implemented
================
- `create_question`: Add an question
```
Usage:
   create_question 1 “How many letters are there in the English alphabet?” 26
```

- `delete_question`: Remove an question
```
Usage:
   delete_question 1
```

- `question`: get all questions data
```
Usage:
   questions
```

- `question`: Find a question based on its number
```
Usage:
   question 1
```

- `answer_question`: check the answer from a question
```
Usage:
   answer_question 1 26
```

#### NOTE
test function is inside the build setup