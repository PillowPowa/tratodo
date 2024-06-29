<script lang="ts">
  import { genId } from "../../lib";
  import type { HTMLInputAttributes } from "svelte/elements";
  import Input from "./Input.svelte";

  type $$Props = HTMLInputAttributes & {
    id: string;
  };

  export let id: $$Props["id"] = "";
  export let value: $$Props["value"] = undefined;
  const htmlFor: string = `${id}-${genId()}`;
</script>

<div class="space-y-1">
  <label class="leading-[110%] text-foreground/80" for={htmlFor}>
    <slot>{$$props.placeholder}</slot>
  </label>
  <Input
    bind:value
    class="w-full rounded-md px-4 py-2"
    id={htmlFor}
    type="text"
    placeholder="Username or email"
    required
    {...$$restProps}
  />
</div>
