<script>
  import { user } from "../../api";
  import Button from "../Button.svelte";
  import SignOutIcon from "../icons/SignOutIcon.svelte";

  const getMePromise = user.getMe();
</script>

<div class="w-full bg-secondary rounded-lg px-2 py-4 backdrop-blur-sm">
  {#await getMePromise}
    <div class="grid grid-cols-[max-content,auto] items-center gap-x-2 gap-y-1">
      <div class="size-10 rounded-full skeleton row-span-2" />
      <div
        style={`width: ${Math.random() * 120}px`}
        class="h-[16px] skeleton min-w-[80px]"
      />
      <div
        style={`width: ${Math.random() * 180}px`}
        class="h-[12px] skeleton min-w-[120px]"
      />
    </div>
  {:then data}
    <div class="flex items-center gap-x-2">
      <div class="size-10 rounded-full bg-primary-gradient row-span-2" />
      <div class="leading-[110%]">
        <span class="text-sm">{data.username}</span>
        <p class="text-xs text-foreground/70">{data.email}</p>
      </div>
      <div class="ml-auto">
        <Button variant="secondary" size="icon">
          <SignOutIcon />
        </Button>
      </div>
    </div>
  {/await}
</div>
