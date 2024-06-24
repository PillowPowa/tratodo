import { $api } from "../lib/ky";
import type { CreateTodoInput, Todo } from "../types/todo";

export function getTodos() {
  return $api.get("todo/").json<Todo[]>();
}

export function getById(id: number) {
  return $api.get(`todo/${id}`).json<Todo>();
}

export function createTodo(input: CreateTodoInput) {
  return $api.post("todo/", { json: input }).json<Todo>();
}

export function deleteTodo() {
  return $api.delete("todo/").json<Todo>();
}