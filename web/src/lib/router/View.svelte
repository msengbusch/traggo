<script lang="ts">
  import type {Component, Route} from "./route";
  import type {Lazy} from "../lazy";

    interface Properties {
      routes: Route[];
      depth: number;
    }

    let { routes, depth = 0 } = $props<Properties>();

    let hasChild = depth < routes.length - 1;
    let nextDepth = depth + 1;

    let route = routes[depth];
    let componentOrPromise: Component | Promise<Component>

    if((route.component as Lazy<Component>).lazy) {
      componentOrPromise = (route.component as Lazy<Component>).lazy().then((module) => module.default);
    } else {
      componentOrPromise = route.component
    }
</script>

{#await componentOrPromise}
  <div>Loading...</div>
{:then component}
  <svelte:component this={component} />
  {#if hasChild}
    <svelte:self depth={nextDepth} routes={routes}/>
  {/if}
{:catch error}
  <div>{error.message}</div>
{/await}
