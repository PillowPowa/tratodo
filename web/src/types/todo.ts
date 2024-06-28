export interface Todo {
  id: number;
  title: string;
  completed: boolean;
}

export interface CreateTodoInput {
  title: string;
}

export type PatchTodoInput = Partial<Omit<Todo, "id">>;

export interface TodoQuery {
  sort_by?: "newest" | "oldest";
  filter?: "completed" | "uncompleted";
}
