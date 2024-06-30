<script lang="ts">
  import { userStore } from "../store/user";
  import { navigate } from "svelte-routing";
  import { get } from "svelte/store";

  const userState = get(userStore);
  if (!userState.user && userState.status === "unauthenticated") {
    navigate("/sign-in");
  } else if (userState.status === "initial") {
    userStore.fetchUser().catch(() => {
      navigate("/sign-in");
    });
  }
</script>

{#if $userStore.status === "initial"}
  <div class="min-h-screen bg-black grid place-items-center"></div>
{:else if $userStore.user}
  <slot />
{/if}
