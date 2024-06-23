export interface User {
  id: number;
  email: string;
  username: string;
}

export interface RegisterUserInput {
  email: string;
  username: string;
  password: string;
}

export interface LoginUserInput {
  login: string;
  password: string;
}
