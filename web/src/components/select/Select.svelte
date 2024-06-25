<script lang="ts">
  import cn from "clsx";

  interface $$Props extends svelteHTML.HTMLAttributes<HTMLDivElement> {}

  import { onDestroy } from "svelte";
  import { getSelectContext, setSelectContext } from "./store";
  setSelectContext();

  export const onValueChange = (newValue: string | undefined) => {};
  const select = getSelectContext();

  const unsubscribe = select.subscribe((selected) => {
    onValueChange(selected.value);
  });

  onDestroy(unsubscribe);
</script>

<div class={cn("relative", $$restProps.class)} {...$$restProps}>
  <slot />
</div>
