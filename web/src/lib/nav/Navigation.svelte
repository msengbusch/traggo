<script>
</script>

{#snippet group({ header, entries })}
  <div class="divider">{header}</div>
  <ul>
    {#each entries as entry}
      <li>
        <button class="btn btn-sm btn-ghost min-w-full justify-start">
          <span class="material-icons">
            {entry.icon}
          </span>
          {entry.name}
        </button>
      </li>
    {/each}
  </ul>
{/snippet}

<div>
  <nav>
    <!-- <label for="sidebar" class="btn btn-primary drawer-button">Open</label> -->
  </nav>
  <aside class="drawer lg:drawer-open">
    <input id="sidebar" type="checkbox" class="drawer-toggle" />
    <div class="drawer-content">
      <slot />
    </div>
    <div class="drawer-side">
      <div class="flex flex-col p-4 w-80 min-h-full bg-base-200">
        <a href="https://github.com/traggo" class="text-3xl self-center text-primary"> traggo </a>
        {@render group({
          header: "Dashboards",
          entries: [
            { name: "Default", icon: "dashboard" },
            { name: "Manager", icon: "list_alt" }
          ]
        })}
        {@render group({
          header: "Timesheet",
          entries: [
            { name: "List", icon: "timeline" },
            { name: "Calendar", icon: "calendar_today" }
          ]
        })}
        {@render group({
          header: "User",
          entries: [
            { name: "Tags", icon: "label" },
            { name: "Devices", icon: "devices" },
            { name: "Settings", icon: "settings" }
          ]
        })}
        {@render group({
          header: "Admin",
          entries: [
            { name: "Users", icon: "supervisor_account" },
            { name: "Settings", icon: "tune" }
          ]
        })}
      </div>
    </div>
  </aside>
</div>
