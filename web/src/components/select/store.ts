import { writable } from "svelte/store";
import { setContext, getContext } from "svelte";

interface SelectedItem {
  value: string;
  label: string;
}

interface SelectStore {
  selected: SelectedItem | null;
  isOpen: boolean;
}

const selectInitialState: SelectStore = Object.freeze({
  selected: null,
  isOpen: false,
});

const createSelectStore = (initialState: SelectStore = selectInitialState) => {
  const store = writable<SelectStore>(initialState);
  const { subscribe, set, update } = store;

  return {
    subscribe,
    set,
    select: (option: SelectedItem) => {
      update((state) => {
        if (state.selected?.value === option.value) {
          return { ...state, selected: null };
        }
        return { ...state, selected: option };
      });
    },
    toggle: () => {
      update((state) => ({ ...state, isOpen: !state.isOpen }));
    },
  };
};

const SELECT_CONTEXT = Symbol("SELECT_CONTEXT");

export const setSelectContext = () => {
  const selectStore = createSelectStore();
  setContext(SELECT_CONTEXT, selectStore);
};

export const getSelectContext = () => {
  return getContext<ReturnType<typeof createSelectStore>>(SELECT_CONTEXT);
};
