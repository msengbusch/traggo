<script lang="ts">

import Layout from "./Layout.svelte";
import {lazy} from "./lib/lazy";
import Router from "./lib/router/Router.svelte";
import type {Route} from "./lib/router/route";

const routes: Route[] = [
  {
    path: "/layout",
    component: Layout,
    children: [
      {
        path: "about",
        component: lazy(() => import("./About.svelte")),
      }
    ]
  },
  {
    path: "/",
    component: lazy(() => import("./Home.svelte")),
  }
];
</script>

<Router {routes}>

</Router>