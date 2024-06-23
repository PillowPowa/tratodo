import { $api } from "../lib/ky";
import type { User } from "../types/user";

export async function getMe(): Promise<User> {
  return $api.get("user/").json<User>();
}
