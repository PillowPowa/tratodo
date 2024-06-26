<script lang="ts">
  import type { HTMLButtonAttributes } from "svelte/elements";
  import cn from "clsx";

  interface $$Props extends HTMLButtonAttributes {
    variant?: keyof typeof buttonVariants;
    size?: keyof typeof buttonSizes;
  }

  const buttonVariants = {
    primary:
      "bg-primary-gradient hover:bg-gradient-to-br rounded-md transition-colors shadow-sm transition-all text-primary-foreground shadow-lg shadow-primary/50",
    secondary:
      "bg-secondary/80 rounded-md hover:bg-secondary shadow-md transition-colors",
    destructive:
      "bg-secondary rounded-md hover:bg-destructive-gradient active:bg-gradient-to-br hover:text-destructive-foreground shadow-md transition-all",
    ghost:
      "bg-transparent hover:bg-secondary/20 rounded-md hover:underline transtion-all",
  };

  const buttonSizes = {
    lg: "px-4 py-2 h-10 font-medium",
    icon: "p-1 size-7 text-sm",
    base: "px-2 py-1 h-7 text-sm",
  };

  export let variant: keyof typeof buttonVariants = "primary";
  export let size: keyof typeof buttonSizes = "base";

  let className = "";

  export { className as class };
</script>

<button
  on:click
  class={cn(
    className,
    "disabled:opacity-80 disabled:cursor-now-allowed",
    buttonVariants[variant],
    buttonSizes[size]
  )}
  {...$$restProps}
>
  <slot />
</button>
