<script lang="ts">
    import type {Tag} from "./Tag";
    import TextBadge from "../component/TextBadge.svelte";
    import ColorSlider from "../component/color/ColorSlider.svelte";

    export let tag: Tag

    let isEditing: boolean = false
</script>

<tr>
    <td>
        {#if isEditing}
            <input type="text" placeholder="name" bind:value={tag.key} class="input input-sm"/>
        {:else}
            {tag.key}
        {/if}
    <td>
        <ColorSlider></ColorSlider>
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
