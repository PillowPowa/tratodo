<script lang="ts">
  import type { Todo } from "../../types/todo";
  import Button from "../Button.svelte";
  import Checkbox from "../Checkbox.svelte";

  function onInput(
    this: HTMLTextAreaElement,
    e: { currentTarget: EventTarget & HTMLTextAreaElement }
  ) {
    todo = { ...todo, title: e.currentTarget.value };
    this.style.height = "auto";
    this.style.height = `${this.scrollHeight}px`;
  }

  export let todo: Partial<Todo> = {};
</script>

<form
  data-completed={todo.completed}
  class="px-2 py-4 bg-white/70 backdrop-blur-sm rounded-md data-[completed='true']:opacity-60 transition-colors"
>
  <div class="flex items-center gap-x-2">
    <Checkbox bind:checked={todo.completed}>
      <textarea
        placeholder="Todo title"
        rows="1"
        value={todo.title ?? ""}
        on:input={onInput}
        class="text-balance max-w-full bg-transparent resize-none text-foreground/80 font-medium"
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
