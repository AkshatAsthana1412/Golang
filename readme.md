## This repo contains a collection of projects that I built using golang to demonstrate and enhance my go concurrency skills.

### Here's a short description about each project:

* **timed_quiz** : A command-line quiz application that reads questions and answers from a CSV file and quizzes the user within a specified time limit. It demonstrates Go’s concurrency patterns by using goroutines and channels to handle user input with timeouts. For each question, a separate goroutine is spawned to read the user’s answer, while a select statement concurrently listens for input or a timeout event. This non-blocking, concurrent design ensures that the program remains responsive and can enforce a time constraint per quiz session efficiently.