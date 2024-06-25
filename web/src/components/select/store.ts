import { writable } from "svelte/store";
import { setContext, getContext } from "svelte";

interface SelectStore {
  value?: string;
  label?: string;
}

const createSelectStore = (initialState: SelectStore) => {
  const store = writable<SelectStore>(initialState);
  const { subscribe, set, update } = store;

  return {
    subscribe,
    set,
    select: (option: SelectStore) => {
      update(() => option);
    },
  };
};

const SELECT_CONTEXT = Symbol("SELECT_CONTEXT");

export const setSelectContext = (initialState: SelectStore = {}) => {
  const selectStore = createSelectStore(initialState);
  setContext(SELECT_CONTEXT, selectStore);
};

export const getSelectContext = () => {
  return getContext<ReturnType<typeof createSelectStore>>(SELECT_CONTEXT);
};
