<script>
  import Navigation from "$lib/nav/Navigation.svelte";
  import "./styles.css";
</script>

<div class="app">
  <Navigation>
    Content
    <slot />
  </Navigation>
</div>
