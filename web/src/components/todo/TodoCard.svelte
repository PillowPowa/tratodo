<script lang="ts">
  import type { Todo } from "src/types/todo";
  import Button from "../Button.svelte";
  import Checkbox from "../Checkbox.svelte";
  import EditIcon from "../icons/EditIcon.svelte";
  import TrashIcon from "../icons/TrashIcon.svelte";
  import { clickOutside } from "../../lib/hooks";
  import cn from "clsx";
  import { onDestroy } from "svelte";

  export let todo: Todo;
  let snapshot: Todo | null = null;

  $: snapshot = editable ? { ...todo } : null;

  const onCancel = () => {
    if (!editable || !snapshot) return;
    todo = snapshot;
    titleRef.textContent = snapshot.title;
    editable = false;
  };

  const onSubmit = () => {
    onCancel();
  };

  const onEdit = () => {
    editable = true;
    setTimeout(() => {
      const selectedText = window.getSelection();
      if (!selectedText) return;
      const selectedRange = document.createRange();
      selectedRange.setStart(
        titleRef.childNodes[0],
        titleRef.textContent?.length ?? 0
      );
      selectedRange.collapse(true);
      selectedText.removeAllRanges();
      selectedText.addRange(selectedRange);
      titleRef.focus();
    });
  };

  const keyDownHandler = (e: KeyboardEvent) => {
    switch (e.key) {
      case "Enter":
        onSubmit();
        break;
      case "Escape":
        onCancel();
        break;
    }
  };

  document.addEventListener("keydown", keyDownHandler, true);
  onDestroy(() =>
    document.removeEventListener("keydown", keyDownHandler, true)
  );

  let editable = false;
  let titleRef: HTMLParagraphElement;
</script>

<div
  use:clickOutside={onSubmit}
  data-completed={editable ? false : todo.completed}
  class={cn(
    "px-2 py-4 bg-white/70 backdrop-blur-sm rounded-md data-[completed='true']:opacity-60 transition-all z-20",
    editable && "shadow-md scale-[1.03]"
  )}
>
  <div class="flex items-center gap-x-2">
    <Checkbox checked={todo.completed} disabled={editable} />
    <div class="text-foreground/80 font-medium">
      <p
        bind:this={titleRef}
        class="text-balance line-clamp-5 focus:outline-none cursor-pointer"
        contenteditable={editable}
      >
        {#if todo.completed && !editable}
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

    {#if editable}
      <Button class="ml-auto" variant="ghost" size="base" on:click={onCancel}>
        Cancel
      </Button>
    {:else}
      <div class="ml-auto flex items-center gap-x-2">
        <Button
          on:click={onEdit}
          aria-label="Edit todo"
          variant="secondary"
          size="icon"
        >
          <EditIcon />
        </Button>

        <Button aria-label="Delete todo" variant="destructive" size="icon">
          <TrashIcon />
        </Button>
      </div>
    {/if}
  </div>
</div>
