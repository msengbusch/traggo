<script lang="ts">
  import View from "./View.svelte";
  import type {Route} from "./route";
  import {resolve} from "./parse";

    interface Properties {
      routes: Route[]
    }

    let { routes } = $props<Properties>();

    let resolved = $state(getRoute(""));

    function getRoute(path: string): Route[] {
      let found = resolve(routes, path);
      if(found) {
        return found;
      }

      return [];
    }

    $effect(() => {
      window.addEventListener("hashchange", () => {
        resolved = getRoute(window.location.hash.slice(1));
      })
    })
</script>

<View routes={resolved} depth={0} />