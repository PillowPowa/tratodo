<script lang="ts">
  import type { HTMLFormAttributes } from "svelte/elements";
  import type { Todo } from "../../types/todo";
  import Button from "../ui/Button.svelte";
  import Checkbox from "../ui/Checkbox.svelte";

  type $$Props = HTMLFormAttributes & {
    todo?: Partial<Todo>;
    "on:change"?: (checked: boolean) => void;
  };

  function onInput(this: HTMLTextAreaElement) {
    this.style.height = "auto";
    this.style.height = `${this.scrollHeight}px`;
  }

  export let todo: Partial<Todo> = {};
</script>

<form
  data-completed={todo.completed}
  class="px-2 py-4 bg-white/70 backdrop-blur-sm rounded-md data-[completed='true']:opacity-60 transition-colors"
  on:submit|preventDefault
  {...$$restProps}
>
  <div class="flex items-center gap-x-2">
    <Checkbox bind:checked={todo.completed}>
      <textarea
        name="title"
        placeholder="Todo title"
        rows="1"
        bind:value={todo.title}
        on:input={onInput}
        class="text-balance focus:outline-none max-w-full bg-transparent resize-none text-foreground/80 font-medium"
        wrap="hard"
        required
        maxlength="255"
      />
    </Checkbox>

    <div class="ml-auto flex items-center gap-x-2">
      <Button type="submit" variant="primary">
        {todo.id ? "Save" : "Add"}
      </Button>
    </div>
  </div>
</form>
