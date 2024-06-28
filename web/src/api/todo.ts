import { $api } from "../lib/ky";
import type {
  CreateTodoInput,
  PatchTodoInput,
  Todo,
  TodoQuery,
} from "../types/todo";

export function getTodos(query: TodoQuery = {}) {
  return $api
    .get("todo/", { searchParams: Object.entries(query) })
    .json<Todo[]>();
}

export function getById(id: number) {
  return $api.get(`todo/${id}`).json<Todo>();
}

export function createTodo(input: CreateTodoInput) {
  return $api.post("todo/", { json: input }).json<Todo>();
}

export function deleteTodo(id: number) {
  return $api.delete(`todo/${id}`).json<Todo>();
}

export function patchTodo(id: number, input: PatchTodoInput) {
  return $api.patch(`todo/${id}`, { json: input }).json<Todo>();
}
