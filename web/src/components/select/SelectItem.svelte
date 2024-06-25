<script lang="ts">
  import { getSelectContext } from "./store";
  import cn from "clsx";

  const select = getSelectContext();

  function selection() {
    select.select({ label, value });
  }

  interface $$Props {
    label: string;
    value: string;
  }

  export let label = "";
  export let value = "";
</script>

<button
  on:click={selection}
  class={cn(
    "block w-full rounded-md h-7 text-sm bg-secondary/70 hover:bg-secondary transition-colors border",
    "text-left px-2 py-1",
    $$restProps.class
  )}
  {...$$restProps}
>
  {label}
</button>
