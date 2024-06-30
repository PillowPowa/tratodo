<script lang="ts">
  import type { RegisterUserInput } from "src/types/user";
  import Button from "../components/ui/Button.svelte";
  import DetailedInput from "../components/ui/DetailedInput.svelte";
  import { FetchError } from "../lib/ky";
  import { auth } from "../api";
  import { Link, navigate } from "svelte-routing";

  let fields: RegisterUserInput = {
    username: "",
    email: "",
    password: "",
  };

  const onSubmit = (e: SubmitEvent) => {
    auth
      .register(fields)
      .then(() => {
        navigate("/");
        if (e.target instanceof HTMLFormElement) {
          e.target?.reset();
        }
      })
      .catch((err) => {
        if (err instanceof FetchError) {
          return alert(err);
        }
        alert(err.message);
      });
  };
</script>

<main class="min-h-screen grid place-items-center">
  <div class="space-y-4 bg-muted max-w-md w-full h-fit rounded-md">
    <h2
      class="text-6xl my-6 uppercase leading-[110%] font-bold text-center text-secondary"
    >
      Sign Up
    </h2>

    <form
      on:submit|preventDefault={onSubmit}
      class="py-6 sm:px-6 px-2 space-y-6 bg-secondary rounded-md"
    >
      <div class="text-center w-full">
        <h2 class="font-semibold text-3xl">Welcome in Tratodo!</h2>
        <p class="text-foreground/80">
          Create an account by filling the form below
        </p>
      </div>
      <div class="space-y-4">
        <DetailedInput
          bind:value={fields.username}
          id="username"
          placeholder="Your username"
          required
        >
          Username
        </DetailedInput>

        <DetailedInput
          bind:value={fields.email}
          id="email"
          placeholder="Your email address"
          type="email"
          required
        >
          Email
        </DetailedInput>

        <DetailedInput
          bind:value={fields.password}
          id="password"
          placeholder="Your password"
          type="password"
          required
        >
          Password
        </DetailedInput>
      </div>
      <div class="space-y-2">
        <Button class="w-full" type="submit" size="lg">Sign Up</Button>
        <p class="text-foreground/80">
          Already have an account?
          <Link
            class="underline text-primary/70 hover:text-primary transition-colors"
            to="/sign-in"
          >
            Sign in!
          </Link>
        </p>
      </div>
    </form>
  </div>
</main>
