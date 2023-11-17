<script lang="ts">
    import {onDestroy, onMount} from "svelte";
    import Dashboard from "./dashboard/Dashboard.svelte";
    import Settings from "./setting/Settings.svelte";
    import Tags from "./tag/Tags.svelte";
    import Page from "./layout/Page.svelte";

    let eventListener;

    // Listen for hash changed event
    onMount(() => {
        eventListener = () => {
            // Get the hash from the URL
            const hash = window.location.hash;
            console.log(hash)

            let component;
            const componentName = hash.replace("#/", "");

            if (componentName === "dashboard") {
                component = Dashboard
            } else if (componentName === "tags") {
                component = Tags
            } else if (componentName === "settings") {
                component = Settings
            } else {
                throw new Error("Component not found")
            }

            const router = document.getElementById("router");
            router.innerHTML = "";

            new component({
                target: router
            });
        }

        window.addEventListener("hashchange", eventListener);
    });

    onDestroy(() => {
        window.removeEventListener("hashchange", eventListener);
    });
</script>

<style>
    @tailwind base;
    @tailwind components;
    @tailwind utilities;
</style>

<Page>
    <Tags></Tags>
</Page>

<!--<Router>-->
<!--    <Route path="*"><Dashboard></Dashboard></Route>-->
<!--    <Route path="/tags"><Tags></Tags></Route>-->
<!--    <Route path="/settings"><Settings></Settings></Route>-->
<!--</Router>-->
