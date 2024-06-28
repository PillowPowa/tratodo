import { writable } from "svelte/store";

function createQueryStore() {
  const queryParams = writable(new URLSearchParams(location.search));

  queryParams.subscribe((params) => {
    let stringifyParams = params.toString();
    if (stringifyParams.length > 1) stringifyParams = "?" + stringifyParams;
    history.replaceState(null, "", "/" + stringifyParams);
  });

  return {
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
