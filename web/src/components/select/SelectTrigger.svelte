<script lang="ts">
  import type { HTMLInputAttributes } from "svelte/elements";
  import cn from "clsx";
  import { getSelectContext } from "./store";

  const selectState = getSelectContext();

  interface $$Props extends HTMLInputAttributes {
    placeholder?: string;
  }

  export let placeholder = "Select an value";
</script>

<slot>
  <input
    {...$$restProps}
    type="button"
    value={$selectState.label ?? placeholder}
    class={cn(
      "w-full rounded-md bg-background border h-10 px-4 py-2 shadow-sm text-left",
      !$selectState.value && "text-foreground/80",
      $$restProps.class
    )}
  />
</slot>
