<script lang="ts">
  import cn from "clsx";
  import { getSelectContext } from "./store";
  interface $$Props extends svelteHTML.HTMLAttributes<HTMLDivElement> {}

  const select = getSelectContext();
</script>

<div
  data-open={$select.isOpen}
  class={cn(
    "w-full min-w-fit z-20 rounded-md bg-background/90 left-0 top-10 translate-y-1 p-1 space-y-1",
    "absolute z-10 left-0 shadow-md backdrop-blur-sm data-[open='false']:hidden overflow-hidden",
    $$restProps.class
  )}
  {...$$restProps}
>
  <slot></slot>
</div>
