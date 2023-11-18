<script lang="ts">

</script>

<style>
    .range-color {
        @apply bg-inherit;
    }

    .range-color::-moz-range-track,
    .range-color::-webkit-slider-runnable-track {
        @apply h-3 rounded-full;

        background: linear-gradient(to right, #ff0000 0%, #ffff00 17%, #00ff00 33%, #00ffff 50%, #0000ff 67%, #ff00ff 83%, #ff0000 100%);
        margin: calc(theme(spacing[3.5] / 2));
    }

    .range-color::-moz-range-thumb,
    .range-color::-webkit-slider-thumb {
        @apply border-0 w-3.5 h-3.5;
    }
</style>

<input type="range" min="0" max="100" value="40" class="range-color"/>