import { getContext, setContext } from "svelte";
import { writable } from "svelte/store";

function createOverlayStore() {
  const store = writable<boolean>(false);
  return {
    toggle: () => store.update((prev) => !prev),
    ...store,
  };
}

export const overlayStore = createOverlayStore();
const CONTEXT = "OVERLAY_CONTEXT";

export function createOverlayContext() {
  const ctx = createOverlayStore();
  setContext(CONTEXT, ctx);
  return ctx;
}

export function getOverlayContext() {
  return getContext<ReturnType<typeof createOverlayStore>>(CONTEXT);
}
