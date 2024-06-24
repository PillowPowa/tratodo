<script lang="ts">
  import { genId } from "../lib";

  const id = genId('checkbox');
  export let checked = false;
</script>

<label for={id} class="flex items-center">
  <input on:click={() => checked = !checked} id={id} type="checkbox" class="hidden" />
  <span
    data-checked={checked}
    class={`
      p-1 size-7 border bg-secondary rounded-md data-[checked='true']:bg-primary mr-2
      cursor-pointer transition-colors group

    `}
  >
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 256 256">
      <rect width="256" height="256" fill="none"/>
      <polyline class="transition-all delay-100 duration-400 stroke-foreground group-data-[checked='true']:stroke-background" points="40 144 96 200 224 72" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="24" stroke-dasharray="270" stroke-dashoffset={checked ? "0" : "270"}/>
    </svg>
  </span>
  <slot />
</label>

<style>
  .test {
    padding: 2rem;
  }
</style>