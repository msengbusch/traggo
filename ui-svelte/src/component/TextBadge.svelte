<script lang="ts">
    import bestContrast from 'get-best-contrast-color'

    export let color: string

    // TODO: Use css color-contrast
    const textColor = bestContrast(color, ['#fff', '#000'])
</script>

<span class="badge badge-lg" style="background-color: {color}; color: {textColor}">
    <slot></slot>
</span>