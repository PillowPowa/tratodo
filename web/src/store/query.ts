import { writable } from "svelte/store";

function createQueryStore() {
  const queryParams = writable(new URLSearchParams(location.search));

  queryParams.subscribe((params) => {
    history.replaceState(null, "", `?${params.toString()}`);
  });

  window.addEventListener("popstate", () => {
    queryParams.set(new URLSearchParams(location.search));
  });

  return {
    ...queryParams,
    subscribe: queryParams.subscribe,
    set: (params: URLSearchParams) => {
      queryParams.set(params);
    },
    update: (key: string, value: string) => {
      const params = new URLSearchParams(location.search);
      params.set(key, value);
      queryParams.set(params);
    },
    delete: (key: string) => {
      const params = new URLSearchParams(location.search);
      params.delete(key);
      queryParams.set(params);
    },
  };
}

export const queryStore = createQueryStore();
