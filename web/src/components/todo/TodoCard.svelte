<script lang="ts">
  import Button from "../ui/Button.svelte";
  import Checkbox from "../ui/Checkbox.svelte";
  import EditIcon from "../icons/EditIcon.svelte";
  import TrashIcon from "../icons/TrashIcon.svelte";
  import { createOverlayContext, Overlay } from "../ui/overlay";
  import { clickOutside } from "../../lib/hooks";
  import { contentEditableMacroFocus } from "../../lib/document";
  import { memento } from "../../lib/memento";

  import cn from "clsx";

  import type { Todo } from "src/types/todo";
  import { todoStore } from "../../store/todo";

  export let todo: Todo;
  const todoMemento = memento<Todo>();
  let overlayStore = createOverlayContext();
  let editable = false;
  let titleRef: HTMLParagraphElement;

  const onDelete = () => {
    todoStore.deleteTodo(todo.id).catch(() => null);
  };

  const toggleTodo = (forceCompleted: boolean = todo.completed) => {
    todoStore
      .updateTodo(todo.id, { completed: forceCompleted })
      .catch(() => null);
  };

  const onEditionSubmit = async () => {
    if (!editable || todoMemento.isSame(todo)) {
      return onEditionCancel();
    }

    todoStore
      .updateTodo(todo.id, todo)
      .then(escapeEditMode)
      .catch(onEditionCancel);
  };

  const onEditionCancel = () => {
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

  // TEMP: hard code, better to use independent component with form
  const keyDownEventHandler = (e: KeyboardEvent) => {
    switch (e.key) {
      case "Enter":
        onEditionSubmit();
        break;
      case "Escape":
        onEditionCancel();
        break;
    }
  };
</script>

<svelte:document on:keydown={keyDownEventHandler} />

<Overlay>
  <div
    use:clickOutside={onEditionCancel}
    data-completed={todo.completed && !$overlayStore}
    class={cn(
      "relative px-2 py-4 bg-white/70 backdrop-blur-sm rounded-md data-[completed='true']:opacity-60 transition-all",
      $overlayStore && "scale-[1.03] rounded-br-none !bg-white z-50"
    )}
  >
    <div class="flex items-center gap-x-2">
      <Checkbox
        on:change={(e) => toggleTodo(e.detail)}
        checked={todo.completed}
        disabled={editable}
      />
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
        <div class="ml-auto flex items-center gap-x-2 animate-fade-in">
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
            on:click={onDelete}
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
          on:click={onEditionSubmit}
          class="absolute rounded-b-md bg-white h-8 px-6 py-1 text-sm right-0 bottom-0 translate-y-[calc(100%-1px)]"
        >
          Save
        </button>
      {/if}
    </div>
  </div>
</Overlay>
