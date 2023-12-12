<script lang="ts">
    import type {ComponentType} from "svelte";

    interface Properties {
      component: ComponentType<any> | undefined
    }

    //let { component } = $props<Properties>();
</script>

<div>
    Layout
    <slot></slot>
</div>