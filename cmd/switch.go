package cmd

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/erwinhermanto31/quiz_master/action"
	"github.com/erwinhermanto31/quiz_master/entity"
	"github.com/gammban/numtow"
	"github.com/gammban/numtow/lang"
)

// NewSwitch creates a new instance of command Switch
func NewSwitch(command string) Switch {
	s := Switch{}
	s.commands = map[string]func() func(string) error{
		"create_question": s.create,
		"update_question": s.update,
		"delete_question": s.delete,
		"question":        s.fetchOne,
		"questions":       s.fetchAll,
		"answer_question": s.answerQuestion,
	}
	return s
}

// Switch represents CLI command switch
type Switch struct {
	commands map[string]func() func(string) error
}

// Switch analyses the CLI args and executes the given command
func (s Switch) Switch(command string) error {
	cmds := strings.Split(command, " ")
	cmdName := cmds[0]
	cmd, ok := s.commands[cmdName]
	if !ok {
		return fmt.Errorf("invalid command '%s'", cmdName)
	}

	return cmd()(command)
}

// create represents the create command which create question
func (s Switch) create() func(string) error {
	return func(command string) error {
		commands := strings.Split(command, " ")
		if err := s.checkArgs(3, commands); err != nil {
			return err
		}

		no, _ := strconv.Atoi(commands[1])
		answers := []string{}
		for i := 2; i < len(commands)-1; i++ {
			answers = append(answers, commands[i])
		}

		answer := strings.Join(answers, " ")

		req := entity.Question{
			No:       no,
			Question: answer,
			Answer:   commands[len(commands)-1],
		}

		err := action.NewCreateQuestion().Handler(context.Background(), req)
		if err != nil {
			return err
		}

		fmt.Printf("Question no %v created: \n", no)
		fmt.Printf("Q: %v \n", answer)
		fmt.Printf("A: %v \n", commands[len(commands)-1])

		fmt.Println(" ")
		return nil
	}
}

// update represents the create command which update question
func (s Switch) update() func(string) error {
	return func(command string) error {

		commands := strings.Split(command, " ")
		if err := s.checkArgs(3, commands); err != nil {
			return err
		}

		no, _ := strconv.Atoi(commands[1])
		answers := []string{}
		for i := 2; i < len(commands)-1; i++ {
			answers = append(answers, commands[i])
		}

		answer := strings.Join(answers, " ")

		req := entity.Question{
			No:       no,
			Question: answer,
			Answer:   commands[len(commands)-1],
		}

		err := action.NewUpdateQuestion().Handler(context.Background(), req)
		if err != nil {
			return err
		}

		fmt.Printf("Question no %v updated: \n", no)
		fmt.Printf("Q: %v \n", answer)
		fmt.Printf("A: %v \n", commands[len(commands)-1])

		fmt.Println(" ")
		return nil
	}
}

// delete represents the create command which delete question
func (s Switch) delete() func(string) error {
	return func(command string) error {

		commands := strings.Split(command, " ")
		if err := s.checkArgs(1, commands); err != nil {
			return err
		}

		no, _ := strconv.Atoi(commands[1])

		req := entity.Question{
			No: no,
		}

		err := action.NewUpdateQuestion().Handler(context.Background(), req)
		if err != nil {
			return err
		}

		fmt.Printf("Question no %v was Deleted! \n", no)

		fmt.Println(" ")
		return nil
	}
}

// Fetch All represents the create command which delete question
func (s Switch) fetchAll() func(string) error {
	return func(command string) error {

		commands := strings.Split(command, " ")
		if err := s.checkArgs(0, commands); err != nil {
			return err
		}

		req := entity.Question{}

		question, err := action.NewGetAllQuestion().Handler(context.Background(), req)
		if err != nil {
			return err
		}

		fmt.Printf("No | Question | Answer \n")
		for _, v := range question {
			fmt.Printf("%v    %v   %v  \n", v.No, v.Question, v.Answer)
		}

		fmt.Println(" ")
		return nil
	}
}

// Fetch one represents the create command which delete question
func (s Switch) fetchOne() func(string) error {
	return func(command string) error {

		commands := strings.Split(command, " ")
		if err := s.checkArgs(1, commands); err != nil {
			return err
		}

		no, _ := strconv.Atoi(commands[1])

		req := entity.Question{
			No: no,
		}

		question, err := action.NewGetQuestion().Handler(context.Background(), req)
		if err != nil {
			return err
		}

		fmt.Printf("Q: %v \n", question.Question)
		fmt.Printf("A: %v \n", question.Answer)

		fmt.Println(" ")
		return nil
	}
}

// answer question represents the create command which answer question
func (s Switch) answerQuestion() func(string) error {
	return func(command string) error {

		commands := strings.Split(command, " ")
		if err := s.checkArgs(1, commands); err != nil {
			return err
		}

		no, _ := strconv.Atoi(commands[1])
		answer := commands[2]
		req := entity.Question{
			No: no,
		}

		question, err := action.NewGetQuestion().Handler(context.Background(), req)
		if err != nil {
			return err
		}

		if answer == question.Answer {
			fmt.Println("Correct!")
		} else {
			answerWord := numtow.MustString(question.Answer, lang.EN)
			if answerWord == answer {
				fmt.Println("Correct!")
			} else {
				fmt.Println("InCorrect!")
			}
		}

		fmt.Println(" ")
		return nil
	}
}

// checkArgs checks if the number of passed in args is greater or equal to min args
func (s Switch) checkArgs(minArgs int, command []string) error {
	if len(command) == 2 && command[1] == "help" {
		return nil
	}
	if len(command)-1 < minArgs {
		fmt.Printf(
			"incorect use of %s help\n",
			command[0],
		)
		return fmt.Errorf(
			"expects at least: %d arg(s), %d provided",
			minArgs, len(command)-2,
		)
	}
	return nil
}
