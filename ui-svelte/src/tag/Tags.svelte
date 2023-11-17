<script lang="ts">
    import {Tag} from "./Tag";
    import TextBadge from "../component/TextBadge.svelte";
    import TagColumn from "./TagColumn.svelte";

    const tags: Tag[] = [
        {
            key: "task",
            color: "#4340bf",
            usages: 9
        },
        {
            key: "project",
            color: "#b3dde6",
            usages: 6
        },
        {
            key: "bug",
            color: "#f87171",
            usages: 3
        }
    ]

    let isEditing = false
</script>

<div>
    <div class="card bg-base-200 mx-auto max-w-7xl">
        <div class="card-body">
                <button class="btn btn-neutral btn-block">Create Tag</button>

                <table class="table table-md">
                    <thead>
                    <tr>
                        <th>Key</th>
                        <th>Color</th>
                        <th class="text-end">Usages</th>
                        <th class="w-1/6 text-end"></th>
                    </tr>
                    </thead>
                    <tbody>
                        {#each tags as tag}
                            <TagColumn {tag} />
                        {/each}
                    </tbody>
                </table>
        </div>
    </div>
</div>