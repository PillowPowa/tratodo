<script lang="ts">
  import cn from "clsx";
  interface $$Props extends svelteHTML.HTMLAttributes<HTMLDivElement> {}
</script>

<div
  class={cn(
    "w-full min-w-fit z-20 rounded-md bg-background/90 left-0 top-10 translate-y-1 p-1 space-y-1",
    "absolute z-10 left-0 shadow-md backdrop-blur-sm",
    $$restProps.class
  )}
  {...$$restProps}
>
  <slot></slot>
</div>
