<script lang="ts">
  import { onMount } from "svelte";
  import { todo, user } from "./api";
  import TodoCard from "./components/todo/TodoCard.svelte";
  import UserCard from "./components/user/UserCard.svelte";
  import Select from "./components/select/Select.svelte";
  import SelectTrigger from "./components/select/SelectTrigger.svelte";
  import SelectContent from "./components/select/SelectContent.svelte";
  import SelectItem from "./components/select/SelectItem.svelte";
  import EditableTodoCard from "./components/todo/EditableTodoCard.svelte";

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

  let todosPromise = todo.getTodos();

  const onAddTodo = (e: SubmitEvent) => {
    e.preventDefault();
    const target = e.target;
    if (!(target instanceof HTMLFormElement)) return;
    const formData = new FormData(target);
    const title = formData.get("title") as string;

    todo
      .createTodo({ title })
      .then(() => {
        todosPromise = todo.getTodos();
        target.reset();
      })
      .catch((e) => alert(JSON.stringify(e, null, 2)));
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

      <Select>
        <SelectTrigger class="min-w-[160px] w-full" placeholder="All" />
        <SelectContent>
          <SelectItem label="Completed" value="completed" />
          <SelectItem label="Uncompleted" value="uncompleted" />
        </SelectContent>
      </Select>
    </div>
    <div
      class="px-2 py-4 sm:p-6 space-y-4 bg-secondary rounded-lg max-h-[550px] overflow-auto"
    >
      <EditableTodoCard on:submit={onAddTodo} />
      {#await todosPromise}
        <p>Loading...</p>
      {:then todos}
        {#each todos as todo (todo.id)}
          <TodoCard {todo} />
        {/each}
      {/await}
    </div>
  </div>
</main>
