/**
 * Single state memento pattern
 */
export function memento<T extends object>() {
  let state: T | null = null;
  return {
    save: (newState: T): void => {
      state = { ...newState };
    },
    restore: (): T | null => {
      if (!state) return null;
      const restoredState = { ...state };
      return restoredState;
    },
    pop: (): T | null => {
      const poppedState = state;
      state = null;
      return poppedState;
    },
    isSame: (newState: T): boolean => {
      if (!state) return false;
      return JSON.stringify(state) === JSON.stringify(newState);
    },
  };
}
