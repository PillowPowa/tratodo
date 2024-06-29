<script lang="ts">
  import { Route, Router } from "svelte-routing";
  import Home from "./routes/Home.svelte";
  import ProtectedRoute from "./components/ProtectedRoute.svelte";

  export let url = "";
</script>

<Router {url}>
  <ProtectedRoute path="/" component={Home} />
  <Route path="/login"></Route>
</Router>
