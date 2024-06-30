<script lang="ts">
  import { Route, Router } from "svelte-routing";
  import WithAuth from "./components/WithAuth.svelte";
  import Home from "./routes/Home.svelte";
  import SignIn from "./routes/SignIn.svelte";
  import SignUp from "./routes/SignUp.svelte";

  export let url = "";
</script>

<Router {url}>
  <Route path="/">
    <WithAuth>
      <Home />
    </WithAuth>
  </Route>
  <Route path="/sign-in"><SignIn /></Route>
  <Route path="/sign-up"><SignUp /></Route>
</Router>
