<!-- This component is actually bad, cuz the overlay located in the container of slot, not at the global level,
 which can cause layout bugs, This was done to introduce svelte context api, once i fully rewrite it into the store api -->
<script lang="ts">
  import { getOverlayContext } from "./store";
  import cn from "clsx";
  const overlay = getOverlayContext();
</script>

{#if !!$overlay}
  <div
    class={cn("fixed top-0 inset-0 !m-0 !p-0 bg-black/20 z-40 animate-fade-in")}
  />
{/if}

<slot />
