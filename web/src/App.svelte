<script lang="ts">
  import { onMount } from "svelte";
  import { user } from "./api";
  import Checkbox from "./components/Checkbox.svelte";

  onMount(() => {
    user
      .getMe()
      .then((data) => {
        console.log(data);
      })
      .catch((e) => {
        console.log(e.json());
      });
  });

  const todosMock = [
    {
      id: 1,
      title: "Learn SvelteKit",
      completed: false,
    },
    {
      id: 2,
      title: "Learn GraphQL",
      completed: false,
    },
    {
      id: 3,
      title: "Learn TypeScript",
      completed: true,
    },
    {
      id: 4,
      title: "Learn TailwindCSS",
      completed: true,
    },
    {
      id: 5,
      title: "Learn Jest",
      completed: false,
    },
  ];
</script>

<main class="min-h-screen grid place-items-center">
  <div class="space-y-12 bg-muted max-w-2xl w-full h-[450px] rounded-md">
    <h2
      class="text-6xl uppercase leading-[110%] font-bold text-center text-secondary"
    >
      Todo List
    </h2>

    <div class="p-6 space-y-4 bg-secondary rounded-lg">
      {#each todosMock as todo (todo.id)}
        <div
          data-completed={todo.completed}
          class="px-2 py-4 bg-background rounded-md data-[completed='true']:opacity-60 transition-colors"
        >
          <div class="flex items-center gap-x-2">
            <Checkbox checked={todo.completed}>
              <div class="text-foreground/80 font-medium">
                {#if todo.completed}
                  <del>{todo.title}</del>
                {:else}
                  {todo.title}
                {/if}
                <p class="text-xs text-foreground/50">
                  {new Intl.DateTimeFormat("en-US", {
                    year: "numeric",
                    month: "long",
                    day: "numeric",
                  }).format(new Date())}
                </p>
              </div>
            </Checkbox>

            <div class="ml-auto">
              <button
                class="px-2 py-1 bg-primary h-7 text-secondary text-sm rounded-md"
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      {/each}
    </div>
  </div>
</main>
