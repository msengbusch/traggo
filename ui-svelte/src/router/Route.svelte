<script lang="ts">
    import {getContext, SvelteComponent} from "svelte";
    import {Router, ROUTER} from "./router";

    export let path: String
    export let component: SvelteComponent

    const router = getContext<Router>(ROUTER)
    router.on(path, () => {

    })
</script>

<div id="component">

</div>