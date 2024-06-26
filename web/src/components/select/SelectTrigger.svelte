<script lang="ts">
  import type { HTMLInputAttributes } from "svelte/elements";
  import cn from "clsx";
  import { getSelectContext } from "./store";
  import ChevronUpIcon from "../icons/ChevronUpIcon.svelte";

  const select = getSelectContext();

  interface $$Props extends HTMLInputAttributes {
    placeholder?: string;
  }

  export let placeholder = "Select an value";
</script>

<slot>
  <button
    {...$$restProps}
    data-open={$select.isOpen}
    on:click={select.toggle}
    class={cn(
      "w-full rounded-md bg-background border h-10 px-4 py-2 shadow-sm text-left gap-x-2 inline-flex items-center group",
      !$select.selected && "text-foreground/80",
      $$restProps.class
    )}
  >
    {$select.selected?.label ?? placeholder}
    <ChevronUpIcon
      class="ml-auto size-4 fill-foreground transition-transform group-data-[open='true']:rotate-180"
    />
  </button>
</slot>
