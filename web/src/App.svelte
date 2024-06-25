<script lang="ts">
  import { onMount } from "svelte";
  import { user } from "./api";
  import Button from "./components/Button.svelte";
  import TodoCard from "./components/todo/TodoCard.svelte";
  import UserCard from "./components/user/UserCard.svelte";
  import Select from "./components/select/Select.svelte";
  import SelectTrigger from "./components/select/SelectTrigger.svelte";
  import SelectContent from "./components/select/SelectContent.svelte";
  import SelectItem from "./components/select/SelectItem.svelte";

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

<main class="min-h-screen py-12 grid place-items-center">
  <div class="space-y-4 bg-muted max-w-2xl w-full h-fit rounded-md">
    <h2
      class="text-6xl mb-12 uppercase leading-[110%] font-bold text-center text-secondary"
    >
      Todo List
    </h2>

    <div class="flex justify-between gap-x-4 px-2 py-4 bg-secondary rounded-md">
      <Button variant="primary" size="lg">Add Todo</Button>

      <Select>
        <SelectTrigger class="min-w-[160px] w-full" />
        <SelectContent>
          <SelectItem label="123" value="123" />
          <SelectItem label="321" value="321" />
        </SelectContent>
      </Select>
    </div>
    <div class="p-6 space-y-4 bg-secondary rounded-lg">
      {#each todosMock as todo (todo.id)}
        <TodoCard {todo} />
      {/each}
    </div>
    <UserCard />
  </div>
</main>
