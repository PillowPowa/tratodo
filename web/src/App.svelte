<script lang="ts">
  import TodoCard from "./components/todo/TodoCard.svelte";
  import UserCard from "./components/user/UserCard.svelte";
  import Select from "./components/select/Select.svelte";
  import SelectTrigger from "./components/select/SelectTrigger.svelte";
  import SelectContent from "./components/select/SelectContent.svelte";
  import SelectItem from "./components/select/SelectItem.svelte";
  import EditableTodoCard from "./components/todo/EditableTodoCard.svelte";
  import { todoStore } from "./store/todo";
  import { queryStore } from "./store/query";

  const filterKeys = ["completed", "uncompleted"] as const;

  $: {
    const query = Object.fromEntries($queryStore.entries());
    todoStore.fetchTodos(query);
  }

  const onValueChange = (newValue: string | undefined) => {
    if (!newValue) {
      queryStore.delete("filter");
    } else {
      queryStore.update("filter", newValue);
    }
  };

  const onAddTodo = (e: SubmitEvent) => {
    const target = e.target;
    if (!(target instanceof HTMLFormElement)) return;
    const formData = new FormData(target);
    const title = formData.get("title") as string;
    todoStore.addTodo({ title });
    target.reset();
  };
</script>

<main class="min-h-screen grid place-items-center">
  <div class="space-y-4 bg-muted max-w-2xl w-full h-fit rounded-md">
    <h2
      class="text-6xl my-6 uppercase leading-[110%] font-bold text-center text-secondary"
    >
      Todo List
    </h2>

    <div class="flex justify-between gap-x-4 px-2 py-4 bg-secondary rounded-md">
      <UserCard />

      <Select on:select={(e) => onValueChange(e.detail)}>
        <SelectTrigger class="min-w-[160px] w-full" placeholder="All" />
        <SelectContent>
          {#each filterKeys as filterKey}
            <SelectItem value={filterKey} />
          {/each}
        </SelectContent>
      </Select>
    </div>
    <div
      class="px-2 py-4 sm:p-6 space-y-4 bg-secondary rounded-lg max-h-[550px] overflow-auto"
    >
      <EditableTodoCard on:submit={onAddTodo} />
      {#if !$todoStore}
        <p>Loading...</p>
      {:else}
        {#each $todoStore as todo (todo.id)}
          <TodoCard {todo} />
        {/each}
      {/if}
    </div>
  </div>
</main>
