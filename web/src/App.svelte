<script lang="ts">
  import { Route, Router } from "svelte-routing";
  import Home from "./routes/Home.svelte";
  import ProtectedRoute from "./components/ProtectedRoute.svelte";
  import SignIn from "./routes/SignIn.svelte";

  export let url = "";
</script>

<Router {url}>
  <ProtectedRoute path="/" component={Home} />
  <Route path="/sign-in"><SignIn /></Route>
</Router>
