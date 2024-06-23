import { $api } from "../lib/ky";
import type { LoginUserInput, RegisterUserInput } from "../types/user";

export function login(input: LoginUserInput): Promise<boolean> {
  return $api.post("auth/login", { json: input }).json<boolean>();
}

export function register(input: RegisterUserInput): Promise<boolean> {
  return $api.post("auth/register", { json: input }).json<boolean>();
}

export function logout(): Promise<boolean> {
  return $api.post("auth/logout").json<boolean>();
}
