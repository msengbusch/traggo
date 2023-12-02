<script lang="ts">
  let count = $state(0)

  const increment = () => {
    count += 1
  }
</script>

<button class="btn btn-md btn-primary" on:click={increment}>
  count is {count}
</button>
