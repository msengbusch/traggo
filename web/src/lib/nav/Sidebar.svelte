<script>
</script>

<div class="drawer lg:drawer-open">
  <input id="sidebar" type="checkbox" class="drawer-toggle" />
  <div class="drawer-content">
    Hello World!
    <label for="sidebar" class="btn btn-primary drawer-button">Open</label>
    <slot />
  </div>
  <div class="drawer-side">
    <label for="sidebar" aria-label="close sidebar" class="drawer-overlay"></label>
    <div class="p-4 w-80 min-h-full bg-base-200 flex flex-col">
      <a href="https://github.com/traggo" class="text-3xl self-center text-primary"> traggo </a>
      <div class="divider"></div>
      <ul class="">
        <li>
          <a>Item 1</a>
        </li>
        <li>
          <a>Item 2</a>
        </li>
      </ul>
    </div>
  </div>
</div>
