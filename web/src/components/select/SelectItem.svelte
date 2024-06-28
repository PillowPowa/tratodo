<script lang="ts">
  import CheckIcon from "../icons/CheckIcon.svelte";
  import { getSelectContext } from "./store";
  import cn from "clsx";

  const select = getSelectContext();

  function selectCurrent() {
    select.select({ label: finalLabel, value });
    select.toggle();
  }

  interface $$Props {
    label?: string;
    value: string;
  }

  export let label = "";
  export let value = "";
  const finalLabel = label || value;
</script>

<button
  on:click={selectCurrent}
  class={cn(
    "flex items-center w-full rounded-sm leading-[100%] hover:bg-secondary/80 transition-colors",
    "text-left px-3 py-2 capitalize",
    $$restProps.class
  )}
  {...$$restProps}
>
  {finalLabel}
  {#if $select.selected?.value === value}
    <CheckIcon class="size-4 ml-auto fill-primary" />
  {/if}
</button>
