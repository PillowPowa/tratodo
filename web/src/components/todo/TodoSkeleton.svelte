<script lang="ts">
  import Checkbox from "../ui/Checkbox.svelte";
  import cn from "clsx";
</script>

<div
  class={cn(
    "px-2 py-4 bg-white/70 backdrop-blur-sm rounded-md data-[completed='true']:opacity-60 transition-all"
  )}
>
  <div class="flex items-center gap-x-2">
    <Checkbox disabled />
    <div class="space-y-1">
      <div
        class="h-[16px] min-w-[120px] skeleton-secondary"
        style={`width: ${Math.random() * 180}px`}
      />
      <div class="h-[12px] w-[64px] skeleton-secondary" />
    </div>

    <div class="ml-auto flex items-center gap-x-2 animate-fade-in">
      <span class="size-7 skeleton-secondary" />
      <span class="size-7 skeleton-secondary" />
    </div>
  </div>
</div>
