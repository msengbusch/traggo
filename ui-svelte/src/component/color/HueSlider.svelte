<script lang="ts">

</script>

<style>
    .range {
        @apply bg-inherit;
    }

    .range::-moz-range-track,
    .range::-webkit-slider-runnable-track {
        @apply h-3 rounded-full;

        background: linear-gradient(to right, #ff0000 0%, #ffff00 17%, #00ff00 33%, #00ffff 50%, #0000ff 67%, #ff00ff 83%, #ff0000 100%);
        margin: calc(theme(spacing[3.5]) / 2);
    }

    .range::-moz-range-thumb,
    .range::-webkit-slider-thumb {
        @apply border-0 w-3.5 h-3.5;
    }
</style>

<input class="range" type="range" min="0" max="100" value="40"/>