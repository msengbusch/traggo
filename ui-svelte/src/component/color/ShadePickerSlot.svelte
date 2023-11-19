<script lang="ts">
    import type { Color } from "../../common/Color";
    import ColorBox from "./ColorBox.svelte";

    let { color, active }: { color: Color, active: boolean } = $props();
</script>

<style>
    
</style>

<div class="w-3 h-3">
    <ColorBox color={color}></ColorBox>
</div>