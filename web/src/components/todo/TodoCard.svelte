<script lang="ts">
  import Button from "../Button.svelte";
  import Checkbox from "../Checkbox.svelte";
  import EditIcon from "../icons/EditIcon.svelte";
  import TrashIcon from "../icons/TrashIcon.svelte";
  import { createOverlayContext, Overlay } from "../overlay";
  import { clickOutside } from "../../lib/hooks";
  import { contentEditableMacroFocus } from "../../lib/document";
  import { memento } from "../../lib/memento";

  import cn from "clsx";

  import type { Todo } from "src/types/todo";
  import { todo as todoActions } from "../../api";

  const todoMemento = memento<Todo>();
  let overlayStore = createOverlayContext();
  let editable = false;
  let titleRef: HTMLParagraphElement;

  export let todo: Todo;

  const onSubmit = async () => {
    if (!editable || todoMemento.isSame(todo)) return;

    todoActions
      .patchTodo(todo.id, { title: todo.title })
      .then(escapeEditMode)
      .catch((e) => alert(JSON.stringify(e, null, 2)));
  };

  const onCancel = () => {
    if (!editable) return;
    const restored = todoMemento.pop();
    if (!restored) return;
    escapeEditMode(restored);
    return;
  };

  const enterEditMode = () => {
    editable = true;
    todoMemento.save(todo);
    overlayStore.toggle();
    contentEditableMacroFocus(titleRef);
  };

  const escapeEditMode = (newTodo: Todo) => {
    editable = false;
    todo = newTodo;
    titleRef.innerText = todo.title;
    overlayStore.toggle();
  };

  const keyDownEventHandler = (e: KeyboardEvent) => {
    switch (e.key) {
      case "Enter":
        onSubmit();
        break;
      case "Escape":
        onCancel();
        break;
    }
  };
</script>

<svelte:document on:keydown={keyDownEventHandler} />

<Overlay>
  <div
    use:clickOutside={onCancel}
    data-completed={todo.completed && !$overlayStore}
    class={cn(
      "relative px-2 py-4 bg-white/70 backdrop-blur-sm rounded-md data-[completed='true']:opacity-60 transition-all",
      $overlayStore && "scale-[1.03] rounded-br-none !bg-white z-50"
    )}
  >
    <div class="flex items-center gap-x-2">
      <Checkbox checked={todo.completed} disabled={editable} />
      <div class="text-foreground/80 font-medium">
        <p
          bind:this={titleRef}
          on:input={(e) => (todo.title = e.currentTarget.textContent ?? "")}
          class={cn(
            "text-balance line-clamp-5 focus:outline-none cursor-pointer rounded-md",
            todo.completed && !editable && "line-through"
          )}
          contenteditable={editable}
        >
          {todo.title}
        </p>
        <p class="text-xs text-foreground/50">
          {new Intl.DateTimeFormat("en-US", {
            year: "numeric",
            month: "long",
            day: "numeric",
          }).format(new Date())}
        </p>
      </div>

      {#if !editable}
        <div class="ml-auto flex items-center gap-x-2">
          <Button
            on:click={enterEditMode}
            aria-label="Edit todo"
            variant="secondary"
            size="icon"
            disabled={editable}
          >
            <EditIcon />
          </Button>

          <Button
            aria-label="Delete todo"
            variant="destructive"
            size="icon"
            disabled={editable}
          >
            <TrashIcon />
          </Button>
        </div>
      {/if}

      {#if editable}
        <button
          on:click={onSubmit}
          class="absolute rounded-b-md bg-white h-8 px-6 py-1 text-sm right-0 bottom-0 translate-y-[calc(100%-1px)]"
        >
          Save
        </button>
      {/if}
    </div>
  </div>
</Overlay>
