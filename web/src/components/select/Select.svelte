<script lang="ts">
  import cn from "clsx";

  interface $$Props extends svelteHTML.HTMLAttributes<HTMLDivElement> {}

  import { createEventDispatcher, onDestroy } from "svelte";
  import { getSelectContext, setSelectContext } from "./store";
  import { derived } from "svelte/store";

  setSelectContext();

  const dispatch = createEventDispatcher<{
    select: string | undefined;
  }>();

  const select = getSelectContext();

  const unsubscribe = derived(
    select,
    ($select) => $select.selected?.value
  ).subscribe((value) => {
    dispatch("select", value);
  });
  onDestroy(unsubscribe);
</script>

<div class={cn("relative", $$restProps.class)} {...$$restProps}>
  <slot />
</div>
