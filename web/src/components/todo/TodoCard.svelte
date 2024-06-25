<script lang="ts">
  import type { Todo } from "src/types/todo";
  import Button from "../Button.svelte";
  import Checkbox from "../Checkbox.svelte";
  import EditIcon from "../icons/EditIcon.svelte";
  import TrashIcon from "../icons/TrashIcon.svelte";

  export let todo: Todo;
</script>

<div
  data-completed={todo.completed}
  class="px-2 py-4 bg-white/70 backdrop-blur-sm rounded-md data-[completed='true']:opacity-60 transition-colors"
>
  <div class="flex items-center gap-x-2">
    <Checkbox checked={todo.completed}>
      <div class="text-foreground/80 font-medium">
        <p class="text-balance line-clamp-5">
          {#if todo.completed}
            <del>{todo.title}</del>
          {:else}
            {todo.title}
          {/if}
        </p>
        <p class="text-xs text-foreground/50">
          {new Intl.DateTimeFormat("en-US", {
            year: "numeric",
            month: "long",
            day: "numeric",
          }).format(new Date())}
        </p>
      </div>
    </Checkbox>

    <div class="ml-auto flex items-center gap-x-2">
      <Button aria-label="Edit todo" variant="secondary" size="icon">
        <EditIcon />
      </Button>

      <Button aria-label="Delete todo" variant="destructive" size="icon">
        <TrashIcon />
      </Button>
    </div>
  </div>
</div>
