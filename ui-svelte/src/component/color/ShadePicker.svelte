<script lang="ts">
    import type { Color } from "../../common/Color";
    import ColorBox from "./ColorBox.svelte";
    import ShadePickerSlot from "./ShadePickerSlot.svelte";

    const color: Color = {
        lightness: 0.8,
        chroma: 0.12,
        hue: 225,
        alpha: 100
    }
</script>

<style>

</style>

<div class="flex">
    <ColorBox color={color}></ColorBox>
    <div class="divider divider-horizontal"></div>
    <div>test2</div>
    <div class="divider divider-horizontal"></div>
    <div>test3</div>
    <div class="divider divider-horizontal"></div>
    <div>test3</div>
    <div class="divider divider-horizontal"></div>
    <div>test3</div>
</div>