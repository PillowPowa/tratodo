import type {
  CreateTodoInput,
  PatchTodoInput,
  Todo,
  TodoQuery,
} from "src/types/todo";
import { todo } from "../api";
import { get, writable } from "svelte/store";

function createTodoStore() {
  const todos = writable<Todo[] | null>(null);

  const fetchTodos = async (query: TodoQuery = {}) => {
    const todosList = await todo.getTodos(query).catch(() => []);
    todos.set(todosList);
  };

  const addTodo = async (input: CreateTodoInput) => {
    const newTodo = await todo.createTodo(input);
    todos.update((t) => (t ? [...t, newTodo] : [newTodo]));
  };

  const deleteTodo = async (id: number) => {
    const initialState = get(todos);
    todos.update((t) => t?.filter((todo) => todo.id !== id) ?? null);
    return todo.deleteTodo(id).catch((err) => {
      todos.set(initialState);
      throw err;
    });
  };

  const updateTodo = async (id: number, input: PatchTodoInput) => {
    const initialState = get(todos);
    todos.update((t) => {
      if (!t) return null;
      return t.map((todo) => {
        if (todo.id === id) {
          return { ...todo, ...input };
        }
        return todo;
      });
    });
    return todo.patchTodo(id, input).catch((err) => {
      todos.set(initialState);
      throw err;
    });
  };

  return {
    ...todos,
    addTodo,
    deleteTodo,
    fetchTodos,
    updateTodo,
  };
}

export const todoStore = createTodoStore();
