<script lang="ts">
    import bestContrast from 'get-best-contrast-color'
    import type { Color } from '../common/Color';

    let { color }: {color: Color} = $props();

    // TODO: Use css color-contrast
    const textColor = bestContrast(color, ['#fff', '#000'])
</script>

<span class="badge badge-lg" style="background-color: {color}; color: {textColor}">
    <slot></slot>
</span>