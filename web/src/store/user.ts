import type { User } from "../types/user";
import { auth, user } from "../api";
import { writable } from "svelte/store";

interface UserStoreState {
  user: User | null;
  status: "unauthenticated" | "authenticated" | "initial";
}

function createUserStore() {
  const state = writable<UserStoreState>({
    user: null,
    status: "initial",
  });

  return {
    subscribe: state.subscribe,
    fetchUser: () => {
      return user
        .getMe()
        .then((user) => {
          state.set({ user, status: "authenticated" });
        })
        .catch((err) => {
          state.set({ user: null, status: "unauthenticated" });
          throw err;
        });
    },
    logout: () => {
      return auth.logout().then(() => {
        state.set({ user: null, status: "unauthenticated" });
      });
    },
  };
}

export const userStore = createUserStore();
