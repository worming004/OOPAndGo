		package main

		type Todo struct {
			Name string
		}

		type ITodoRepository interface {
			GetTodoFromRepo(id int) Todo
		}

		type TodoService struct {
			repository ITodoRepository
		}

		// Ceci est une méthode
			func 
			(todoservice TodoService)
			 GetTodo(id int) Todo { 
				return todoservice
					.repository
					.GetTodoFromRepo(id)
			}

		func main() {

		}
