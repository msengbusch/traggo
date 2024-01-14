<script>
  import "./styles.css";

  import Navigation from "$lib/nav/Navigation.svelte";
</script>

<div class="app">
  <Navigation>
    Content
    <slot />
  </Navigation>
</div>
