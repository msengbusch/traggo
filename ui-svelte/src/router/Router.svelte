<script lang="ts">
    import {ROUTER, Router} from "./router";
    import {setContext} from "svelte";

    const router = new Router()
    setContext(ROUTER, router)
</script>