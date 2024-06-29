<script>
  import { userStore } from "../../store/user";
</script>

<div
  class="sticky top-10 grid grid-cols-[max-content,auto] items-center gap-x-2 gap-y-1"
>
  {#if $userStore.status === "initial"}
    <div class="size-10 rounded-full skeleton row-span-2" />
    <div
      style={`width: ${Math.random() * 120}px`}
      class="h-[16px] skeleton min-w-[80px]"
    />
    <div
      style={`width: ${Math.random() * 180}px`}
      class="h-[12px] skeleton min-w-[120px]"
    />
  {:else if $userStore.user}
    <div class="size-10 rounded-full bg-primary-gradient row-span-2" />
    <div class="leading-[110%]">
      <span class="text-sm">{$userStore.user.username}</span>
      <p class="text-xs text-foreground/70">{$userStore.user.email}</p>
    </div>
  {/if}
</div>
