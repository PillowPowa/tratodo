<script lang="ts">
  import type { HTMLInputAttributes } from "svelte/elements";
  import { genId } from "../lib";
  import cn from "clsx";
  import { createEventDispatcher } from "svelte";

  interface $$Props extends HTMLInputAttributes {
    checked?: boolean;
    disabled?: boolean;
  }

  const dispatch = createEventDispatcher();
  const onCheckedChange = (e: { currentTarget: HTMLInputElement }) => {
    checked = e.currentTarget.checked;
    dispatch("change", checked);
  };

  const id = genId("checkbox");
  export let checked = false;
  export let disabled = false;
</script>

<label for={id} class="flex items-center space-x-2">
  <span
    data-checked={checked}
    class={cn(
      "p-1 size-7 border bg-secondary rounded-md data-[checked='true']:bg-primary-gradient hover:bg-gradient-to-br",
      "cursor-pointer transition-all group",
      disabled && "cursor-not-allowed opacity-70"
    )}
  >
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 256 256">
      <rect width="256" height="256" fill="none" />
      <polyline
        class="transition-all delay-100 duration-400 stroke-foreground group-data-[checked='true']:stroke-background"
        points="40 144 96 200 224 72"
        fill="none"
        stroke="currentColor"
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="24"
        stroke-dasharray="270"
        stroke-dashoffset={checked ? "0" : "270"}
      />
    </svg>
  </span>
  <slot />
  <input
    on:change={onCheckedChange}
    {id}
    {disabled}
    value={checked}
    type="checkbox"
    class="hidden"
    {...$$restProps}
  />
</label>
