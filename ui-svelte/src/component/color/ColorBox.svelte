<script lang="ts">
    import type { Color } from "../../common/Color";

    let { color } : { color: Color } = $props;
</script>

<div style="background-color: oklch({color.lightness} {color.chroma} {color.hue});">

</div>