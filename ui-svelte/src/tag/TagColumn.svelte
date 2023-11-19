<script lang="ts">
    import type {Tag} from "./Tag";
    import TextBadge from "../component/TextBadge.svelte";
    import ColorSlider from "../component/color/HueSlider.svelte";
    import ShadePicker from "../component/color/ShadePicker.svelte";
    import HueSlider from "../component/color/HueSlider.svelte";

    let { tag } = $props();

    let isEditing: boolean = $state(false)
</script>

<tr>
    <td>
        {#if isEditing}
            <input type="text" placeholder="name" bind:value={tag.key} class="input input-sm"/>
        {:else}
            {tag.key}
        {/if}
    <td>
        <HueSlider></HueSlider>
        <ShadePicker></ShadePicker>
        <TextBadge color={tag.color}>{tag.color}</TextBadge>
    </td>
    <td class="text-end">{tag.usages}</td>
    <td class="text-end">
        <button class="btn btn-sm btn-circle" on:click={() => isEditing = true}>
                                        <span class="material-icons">
                                            {#if isEditing}
                                                save
                                            {:else}
                                                edit
                                            {/if}
                                        </span>
        </button>
        <button class="btn btn-sm btn-circle" on:click={() => isEditing = false}>
                                        <span class="material-icons">
                                            {#if isEditing}
                                                close
                                            {:else}
                                                delete
                                            {/if}
                                        </span>
        </button>
    </td>
</tr>
